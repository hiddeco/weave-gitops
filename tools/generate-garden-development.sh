sed 's/{{ .Image }}/${modules.wego-app-image.outputs.deployment-image-id}/g' manifests/wego-app/deployment.yaml > .garden-deployment.yaml