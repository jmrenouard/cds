package cdsclient

import (
	"context"
	"io"
	"time"

	"github.com/ovh/cds/engine/api/worker"
	"github.com/ovh/cds/sdk"
)

// Interface is the main interface for cdsclient package
type Interface interface {
	ActionDelete(actionName string) error
	ActionGet(actionName string, mods ...RequestModifier) (*sdk.Action, error)
	ActionList() ([]sdk.Action, error)
	APIURL() string
	ApplicationCreate(string, *sdk.Application) error
	ApplicationDelete(string, string) error
	ApplicationGet(string, string, ...RequestModifier) (*sdk.Application, error)
	ApplicationList(string) ([]sdk.Application, error)
	ApplicationKeysList(string, string) ([]sdk.ApplicationKey, error)
	ApplicationKeyCreate(string, string, *sdk.ApplicationKey) error
	ApplicationKeysDelete(string, string, string) error
	ApplicationVariablesList(key string, appName string) ([]sdk.Variable, error)
	ApplicationVariableCreate(projectKey string, appName string, variable *sdk.Variable) error
	ApplicationVariableDelete(projectKey string, appName string, variable string) error
	ApplicationVariableUpdate(projectKey string, appName string, variable *sdk.Variable) error
	ConfigUser() (map[string]string, error)
	EnvironmentCreate(string, *sdk.Environment) error
	EnvironmentDelete(string, string) error
	EnvironmentGet(string, string, ...RequestModifier) (*sdk.Environment, error)
	EnvironmentList(string) ([]sdk.Environment, error)
	EnvironmentKeysList(string, string) ([]sdk.EnvironmentKey, error)
	EnvironmentKeyCreate(string, string, *sdk.EnvironmentKey) error
	EnvironmentKeysDelete(string, string, string) error
	EnvironmentVariablesList(key string, envName string) ([]sdk.Variable, error)
	EnvironmentVariableCreate(projectKey string, envName string, variable *sdk.Variable) error
	EnvironmentVariableDelete(projectKey string, envName string, keyName string) error
	EnvironmentVariableUpdate(projectKey string, envName string, variable *sdk.Variable) error
	GroupCreate(group *sdk.Group) error
	GroupDelete(name string) error
	GroupGenerateToken(groupName, expiration string) (*sdk.Token, error)
	GroupGet(name string, mods ...RequestModifier) (*sdk.Group, error)
	GroupList() ([]sdk.Group, error)
	GroupUserAdminSet(groupname string, username string) error
	GroupUserAdminRemove(groupname, username string) error
	GroupUserAdd(groupname string, users []string) error
	GroupUserRemove(groupname, username string) error
	HatcheryRefresh(int64) error
	HatcheryRegister(sdk.Hatchery) (*sdk.Hatchery, bool, error)
	MonStatus() ([]string, error)
	PipelineDelete(projectKey, name string) error
	PipelineExport(projectKey, name string, exportWithPermissions bool, exportFormat string) ([]byte, error)
	PipelineImport(projectKey string, content []byte, format string, force bool) ([]string, error)
	PipelineList(projectKey string) ([]sdk.Pipeline, error)
	ProjectCreate(*sdk.Project) error
	ProjectDelete(string) error
	ProjectGet(string, ...RequestModifier) (*sdk.Project, error)
	ProjectList() ([]sdk.Project, error)
	ProjectKeysList(string) ([]sdk.ProjectKey, error)
	ProjectKeyCreate(string, *sdk.ProjectKey) error
	ProjectKeysDelete(string, string) error
	ProjectVariablesList(key string) ([]sdk.Variable, error)
	ProjectVariableCreate(projectKey string, variable *sdk.Variable) error
	ProjectVariableDelete(projectKey string, variable string) error
	ProjectVariableUpdate(projectKey string, variable *sdk.Variable) error
	Queue() ([]sdk.WorkflowNodeJobRun, []sdk.PipelineBuildJob, error)
	QueuePolling(context.Context, chan<- sdk.WorkflowNodeJobRun, chan<- sdk.PipelineBuildJob, chan<- error, time.Duration, int) error
	QueueTakeJob(sdk.WorkflowNodeJobRun, bool) (*worker.WorkflowNodeJobRunInfo, error)
	QueueJobBook(isWorkflowJob bool, id int64) error
	QueueJobInfo(id int64) (*sdk.WorkflowNodeJobRun, error)
	QueueJobSendSpawnInfo(isWorkflowJob bool, id int64, in []sdk.SpawnInfo) error
	QueueSendResult(int64, sdk.Result) error
	QueueArtifactUpload(id int64, tag, filePath string) error
	Requirements() ([]sdk.Requirement, error)
	ServiceRegister(sdk.Service) (string, error)
	TemplateList() ([]sdk.Template, error)
	TemplateGet(name string) (*sdk.Template, error)
	TemplateApplicationCreate(projectKey, name string, template *sdk.Template) error
	UserLogin(username, password string) (bool, string, error)
	UserList() ([]sdk.User, error)
	UserSignup(username, fullname, email, callback string) error
	UserGet(username string) (*sdk.User, error)
	UserGetGroups(username string) (map[string][]sdk.Group, error)
	UserReset(username, email, callback string) error
	UserConfirm(username, token string) (bool, string, error)
	Version() (*sdk.Version, error)
	WorkerList() ([]sdk.Worker, error)
	WorkerModelSpawnError(id int64, info string) error
	WorkerModelsEnabled() ([]sdk.Model, error)
	WorkerModels() ([]sdk.Model, error)
	WorkerRegister(worker.RegistrationForm) (*sdk.Worker, bool, error)
	WorkerSetStatus(sdk.Status) error
	WorkflowList(projectKey string) ([]sdk.Workflow, error)
	WorkflowGet(projectKey, name string) (*sdk.Workflow, error)
	WorkflowRunGet(projectKey string, name string, number int64) (*sdk.WorkflowRun, error)
	WorkflowRunArtifacts(projectKey string, name string, number int64) ([]sdk.Artifact, error)
	WorkflowRunFromHook(projectKey string, workflowName string, hook sdk.WorkflowNodeRunHookEvent) (*sdk.WorkflowRun, error)
	WorkflowRunFromManual(projectKey string, workflowName string, manual sdk.WorkflowNodeRunManual, number, fromNodeID int64) (*sdk.WorkflowRun, error)
	WorkflowNodeRun(projectKey string, name string, number int64, nodeRunID int64) (*sdk.WorkflowNodeRun, error)
	WorkflowNodeRunArtifacts(projectKey string, name string, number int64, nodeRunID int64) ([]sdk.Artifact, error)
	WorkflowNodeRunArtifactDownload(projectKey string, name string, artifactID int64, w io.Writer) error
	WorkflowNodeRunJobStep(projectKey string, workflowName string, number int64, nodeRunID, job int64, step int) (*sdk.BuildState, error)
	WorkflowNodeRunRelease(projectKey string, workflowName string, runNumber int64, nodeRunID int64, release sdk.WorkflowNodeRunRelease) error
	WorkflowAllHooksList() ([]sdk.WorkflowNodeHook, error)
}
