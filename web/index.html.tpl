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
  <h1>GitOps Runtime versions. Generated at {{.Now.Format "02 Jan 06 15:04 MST"}}. See <a href="https://artifacthub.io/packages/helm/codefresh-gitops-runtime/gitops-runtime">Artifact Hub</a> for more details. Next update in 2 hours.
		</h1>

    {{range $item, $gitHubRelease := .VersionsFound}}

  <div class="tree-container">
    <a href="https://codefresh.io" target="_blank" class="card-link">
      <div class="card codefresh">
        <span class="date-label">Feb 2024</span>
        GitOps Runtime
        <span class="version-label"> {{$gitHubRelease.GitOpsRuntime.Version}}</span>
      </div>
    </a>
    <!-- Arrow from Codefresh to Helm chart -->
    <div class="arrow vertical"></div>
    <a href="https://codefresh.io" target="_blank" class="card-link">
      <div class="card helm-chart">
        <span class="date-label">2 Feb 2024</span>
        Argo Helm chart
        <span class="version-label">version 1.2.3 - v.345</span>
      </div>
    </a>
    <div class="arrow horizontal"></div>
    <!-- Grouped cards -->
    <div class="grouped-cards">
      <a href="https://codefresh.io" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label">Feb 2024</span>
          Argo CD
          <span class="version-label">version 1.2.3 - v.345</span>
        </div>
      </a>
      <a href="https://codefresh.io" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label">Feb 2024</span>
          Argo Rollouts
          <span class="version-label">version 1.2.3 - v.345</span>
        </div>
      </a>
      <a href="https://codefresh.io" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label">Feb 2024</span>
          Argo Workflows
          <span class="version-label">version 1.2.3 - v.345</span>
        </div>
      </a>
      <a href="https://codefresh.io" target="_blank" class="card-link">
        <div class="card">
          <span class="date-label">Feb 2024</span>
          Argo Events
          <span class="version-label">version 1.2.3 - v.345</span>
        </div>
      </a>
    </div>
  </div>
  <hr/>
  {{end}}
</body>
</html>