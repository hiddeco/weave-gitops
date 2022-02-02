import * as React from "react";
import styled from "styled-components";
import { Button } from "../..";
import DataTable from "../../components/DataTable";
import Link from "../../components/Link";
import Page from "../../components/Page";
import { useGetKustomizations } from "../../hooks/kustomizations";
import { V2Routes } from "../../lib/types";
import { formatURL } from "../../lib/utils";

type Props = {
  className?: string;
};

function KustomizationList({ className }: Props) {
  const { data: automations, error, isLoading } = useGetKustomizations();
  return (
    <Page
      error={error}
      loading={isLoading}
      title="Kustomizations"
      actions={
        <Link to={formatURL(V2Routes.AddKustomization)}>
          <Button>Add Kustomization</Button>
        </Link>
      }
      className={className}
    >
      <DataTable
        sortFields={["name"]}
        fields={[{ value: "name", label: "Name" }]}
        rows={automations?.kustomizations}
      />
    </Page>
  );
}

export default styled(KustomizationList).attrs({
  className: KustomizationList.name,
})``;