import _ from "lodash";
import * as React from "react";
import styled from "styled-components";
import { Automation } from "../hooks/automations";
import {
  HelmRelease,
  Kustomization,
  SourceRefSourceKind,
} from "../lib/api/core/types.pb";
import { formatURL } from "../lib/nav";
import { AutomationType, V2Routes } from "../lib/types";
import { statusSortHelper } from "../lib/utils";
import { Field, SortType } from "./DataTable";
import FilterableTable, {
  filterConfigForStatus,
  filterConfigForString,
} from "./FilterableTable";
import KubeStatusIndicator, { computeMessage } from "./KubeStatusIndicator";
import Link from "./Link";
import SourceLink from "./SourceLink";
import Timestamp from "./Timestamp";

type Props = {
  className?: string;
  automations?: Automation[];
  appName?: string;
  hideSource?: boolean;
};

function AutomationsTable({ className, automations, hideSource }: Props) {
  const initialFilterState = {
    ...filterConfigForString(automations, "type"),
    ...filterConfigForString(automations, "namespace"),
    ...filterConfigForString(automations, "clusterName"),
    ...filterConfigForStatus(automations),
  };

  let fields: Field[] = [
    {
      label: "Name",
      value: (k) => {
        const route =
          k.type === AutomationType.Kustomization
            ? V2Routes.Kustomization
            : V2Routes.HelmRelease;
        return (
          <Link
            to={formatURL(route, {
              name: k.name,
              namespace: k.namespace,
              clusterName: k.clusterName,
            })}
          >
            {k.name}
          </Link>
        );
      },
      sortValue: ({ name }) => name,
      textSearchable: true,
    },
    {
      label: "Type",
      value: "type",
    },
    {
      label: "Namespace",
      value: "namespace",
    },
    {
      label: "Cluster",
      value: "clusterName",
    },
    {
      label: "Source",
      value: (a: Automation) => {
        let sourceKind;
        let sourceName;

        if (a.type === AutomationType.Kustomization) {
          sourceKind = a.sourceRef?.kind;
          sourceName = a.sourceRef?.name;
        } else {
          sourceKind = SourceRefSourceKind.HelmChart;
          sourceName = (a as HelmRelease).helmChart.name;
        }

        return (
          <SourceLink
            sourceRef={{
              kind: sourceKind,
              name: sourceName,
              namespace: a.sourceRef?.namespace,
            }}
          />
        );
      },
      sortValue: (a: Automation) => a.sourceRef?.name,
    },
    {
      label: "Status",
      value: (a: Automation) =>
        a.conditions.length > 0 ? (
          <KubeStatusIndicator
            short
            conditions={a.conditions}
            suspended={a.suspended}
          />
        ) : null,
      sortType: SortType.number,
      sortValue: statusSortHelper,
    },
    {
      label: "Message",
      value: (a: Automation) => computeMessage(a.conditions),
      sortValue: ({ conditions }) => computeMessage(conditions),
    },
    {
      label: "Revision",
      value: "lastAttemptedRevision",
    },
    {
      label: "Last Updated",
      value: (a: Automation) => (
        <Timestamp time={(a as Kustomization).lastHandledReconciledAt} />
      ),
    },
  ];

  if (hideSource) fields = _.filter(fields, (f) => f.label !== "Source");

  return (
    <FilterableTable
      fields={fields}
      filters={initialFilterState}
      rows={automations}
      className={className}
    />
  );
}

export default styled(AutomationsTable).attrs({
  className: AutomationsTable.name,
})``;
