{{- define "boundary.name" -}}{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}{{- end -}}
{{- define "boundary.fullname" -}}{{- printf "%s-%s" .Release.Name (include "boundary.name" .) | trunc 63 | trimSuffix "-" -}}{{- end -}}
{{- define "boundary.labels" -}}
app.kubernetes.io/name: {{ include "boundary.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | quote }}
{{- end -}}
