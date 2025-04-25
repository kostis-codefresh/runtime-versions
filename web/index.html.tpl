<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Codefresh runtime versions</title>
  <link rel="shortcut icon" href="favicon.png" type="image/png" />
  <link rel="stylesheet" href="style.css">
</head>
<body>
  <h1>GitOps Runtime versions. Generated at {{.Now.Format "02 Jan 2006 15:04 MST"}}. See <a href="https://artifacthub.io/packages/helm/codefresh-gitops-runtime/gitops-runtime">Artifact Hub</a> for more details. Next update in 2 hours.
		</h1>

    {{range $item, $gitHubRelease := .VersionsFound}}

  <div class="tree-container">
    <a href="{{$gitHubRelease.GitOpsRuntime.Link}}" target="_blank" class="card-link">
      <div class="card codefresh">
        <span class="date-label">{{$gitHubRelease.GitOpsRuntime.Date.Format "02 Jan 2006"}}</span>
        GitOps Runtime
        <span class="version-label"> {{$gitHubRelease.GitOpsRuntime.Version}}</span>
      </div>
    </a>

    <div class="arrow vertical"></div>

    <div class="grouped-cards">
    <a href="{{$gitHubRelease.ArgoCD.ArgoHelmChart.Link}}" target="_blank" class="card-link">
      <div class="card helm-chart">
        <span class="date-label">argo-cd</span>
        Argo Helm chart
        <span class="version-label">{{$gitHubRelease.ArgoCD.ArgoHelmChart.Version}}</span>
      </div>
    </a>
    <a href="{{$gitHubRelease.ArgoRollouts.ArgoHelmChart.Link}}" target="_blank" class="card-link">
      <div class="card helm-chart">
        <span class="date-label">argo-rollouts</span>
        Argo Helm chart
        <span class="version-label">{{$gitHubRelease.ArgoRollouts.ArgoHelmChart.Version}}</span>
      </div>
    </a>
    <a href="{{$gitHubRelease.ArgoWorkflows.ArgoHelmChart.Link}}" target="_blank" class="card-link">
      <div class="card helm-chart">
        <span class="date-label">argo-workflows</span>
        Argo Helm chart
        <span class="version-label">{{$gitHubRelease.ArgoWorkflows.ArgoHelmChart.Version}}</span>
      </div>
    </a>
    <a href="{{$gitHubRelease.ArgoEvents.ArgoHelmChart.Link}}" target="_blank" class="card-link">
      <div class="card helm-chart">
        <span class="date-label">argo-events</span>
        Argo Helm chart
        <span class="version-label">{{$gitHubRelease.ArgoEvents.ArgoHelmChart.Version}}</span>
      </div>
    </a>
    </div>
    <div class="arrow horizontal"></div>
    <!-- Grouped cards -->
    <div class="grouped-cards">
      <a href="{{$gitHubRelease.ArgoCD.SourceCodeRepo.Link}}" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label"></span>
          Argo CD
          <span class="version-label">{{$gitHubRelease.ArgoCD.SourceCodeRepo.Version}}</span>
        </div>
      </a>
      <a href="{{$gitHubRelease.ArgoRollouts.SourceCodeRepo.Link}}" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label"></span>
          Argo Rollouts
          <span class="version-label">{{$gitHubRelease.ArgoRollouts.SourceCodeRepo.Version}}</span>
        </div>
      </a>
      <a href="{{$gitHubRelease.ArgoWorkflows.SourceCodeRepo.Link}}" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label"></span>
          Argo Workflows
          <span class="version-label">{{$gitHubRelease.ArgoWorkflows.SourceCodeRepo.Version}}</span>
        </div>
      </a>
      <a href="{{$gitHubRelease.ArgoEvents.SourceCodeRepo.Link}}" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label"></span>
          Argo Events
          <span class="version-label">{{$gitHubRelease.ArgoEvents.SourceCodeRepo.Version}}</span>
        </div>
      </a>
    </div>
  </div>
  <hr/>
  {{end}}
</body>
</html>