package postman_api

type DeleteApisApiIdRequest struct {
    APIID   string
}

type DeleteApisApiIdSchemasSchemaIdFilesFilePathRequest struct {
    APIID   string
    SchemaID   string
    FilePath   string
}

type DeleteApisApiIdVersionsVersionIdRequest struct {
    APIID   string
    VersionID   string
}

type DeleteCollectionsCollectionIdRequest struct {
    CollectionID   string
}

type DeleteCollectionsCollectionIdFoldersFolderIdRequest struct {
    CollectionID   string
    FolderID   string
}

type DeleteCollectionsCollectionIdRequestsRequestIdRequest struct {
    CollectionID   string
    RequestID   string
}

type DeleteCollectionsCollectionIdResponsesResponseIdRequest struct {
    CollectionID   string
    ResponseID   string
}

type DeleteEnvironmentsEnvironmentIdRequest struct {
    EnvironmentID   string
}

type DeleteMocksMockIdRequest struct {
    MockID   string
}

type DeleteMocksMockIdServerResponsesServerResponseIdRequest struct {
    MockID   string
    ServerResponseID   string
}

type DeleteMocksMockIdUnpublishRequest struct {
    MockID   string
}

type DeleteMonitorsMonitorIdRequest struct {
    MonitorID   string
}

type DeleteNetworkPrivateElementTypeElementIdRequest struct {
    ElementType   string
    ElementID   string
}

type DeleteScimV2GroupsGroupIdRequest struct {
    GroupID   string
}

type DeleteWorkspacesWorkspaceIdRequest struct {
    WorkspaceID   string
}

type GetApisRequest struct {
    WorkspaceID   string
    CreatedBy   *int
    Cursor   *string
    Description   *string
    Limit   *int
}

type GetApisApiIdRequest struct {
    APIID   string
    Include   *[]string
}

type GetApisApiIdCollectionsCollectionIdRequest struct {
    APIID   string
    CollectionID   string
    VersionID   *string
}

type GetApisApiIdSchemasSchemaIdRequest struct {
    APIID   string
    SchemaID   string
    Bundled   *bool
    VersionID   *string
}

type GetApisApiIdSchemasSchemaIdFilesRequest struct {
    APIID   string
    SchemaID   string
    Cursor   *string
    Limit   *int
    VersionID   *string
}

type GetApisApiIdSchemasSchemaIdFilesFilePathRequest struct {
    APIID   string
    SchemaID   string
    FilePath   string
    VersionID   *string
}

type GetApisApiIdTagsRequest struct {
    APIID   string
}

type GetApisApiIdTasksTaskIdRequest struct {
    APIID   string
    TaskID   string
}

type GetApisApiIdVersionsRequest struct {
    APIID   string
    Cursor   *string
    Limit   *int
}

type GetApisApiIdVersionsVersionIdRequest struct {
    APIID   string
    VersionID   string
}

type GetAuditLogsRequest struct {
    Cursor   *string
    Limit   *int
    OrderBy   *string
    Since   *string
    Until   *string
}

type GetCollectionsRequest struct {
    Name   *string
    WorkspaceID   *string
}

type GetCollectionsCollectionIdRequest struct {
    CollectionID   string
    AccessKey   *string
}

type GetCollectionsCollectionIdFoldersFolderIdRequest struct {
    CollectionID   string
    FolderID   string
    IDS   *bool
    Populate   *bool
    Uid   *bool
}

type GetCollectionsCollectionIdRequestsRequestIdRequest struct {
    CollectionID   string
    RequestID   string
    IDS   *string
    Populate   *bool
    Uid   *bool
}

type GetCollectionsCollectionIdResponsesResponseIdRequest struct {
    CollectionID   string
    ResponseID   string
    IDS   *bool
    Populate   *bool
    Uid   *bool
}

type GetCollectionsCollectionIdTagsRequest struct {
    CollectionID   string
}

type GetCollectionsCollectionIdTransformationsRequest struct {
    CollectionID   string
}

type GetDetectedSecretsSecretIdLocationsRequest struct {
    SecretID   string
    WorkspaceID   string
    Cursor   *string
    Limit   *int
}

type GetEnvironmentsRequest struct {
    WorkspaceID   *string
}

type GetEnvironmentsEnvironmentIdRequest struct {
    EnvironmentID   string
}

type GetMocksRequest struct {
    TeamID   *string
    Workspace   *string
}

type GetMocksMockIdRequest struct {
    MockID   string
}

type GetMocksMockIdCallLogsRequest struct {
    MockID   string
    Cursor   *string
    Direction   *string
    Include   *string
    Limit   *float64
    RequestMethod   *string
    RequestPath   *string
    ResponseStatusCode   *float64
    ResponseType   *string
    Since   *string
    Sort   *string
    Until   *string
}

type GetMocksMockIdServerResponsesRequest struct {
    MockID   string
}

type GetMocksMockIdServerResponsesServerResponseIdRequest struct {
    MockID   string
    ServerResponseID   string
}

type GetMonitorsRequest struct {
    Workspace   *string
}

type GetMonitorsMonitorIdRequest struct {
    MonitorID   string
}

type GetNetworkPrivateRequest struct {
    AddedBy   *int
    CreatedBy   *int
    Description   *string
    Direction   *string
    Limit   *int
    Name   *string
    Offset   *int
    ParentFolderID   *int
    Since   *string
    Sort   *string
    Summary   *string
    Type   *string
    Until   *string
}

type GetNetworkPrivateNetworkEntityRequestAllRequest struct {
    Direction   *string
    Limit   *int
    Name   *string
    Offset   *int
    RequestedBy   *int
    Since   *string
    Sort   *string
    Status   *string
    Type   *string
    Until   *string
}

type GetScimV2GroupsRequest struct {
    Count   *float64
    Filter   *string
    StartIndex   *float64
}

type GetScimV2GroupsGroupIdRequest struct {
    GroupID   string
}

type GetScimV2UsersRequest struct {
    Count   *float64
    Filter   *string
    StartIndex   *float64
}

type GetScimV2UsersUserIdRequest struct {
    UserID   string
}

type GetTagsSlugEntitiesRequest struct {
    Slug   string
    Cursor   *string
    Direction   *string
    EntityType   *string
    Limit   *int
}

type GetWorkspacesRequest struct {
    Type   *string
}

type GetWorkspacesWorkspaceIdRequest struct {
    WorkspaceID   string
}

type GetWorkspacesWorkspaceIdGlobalVariablesRequest struct {
    WorkspaceID   string
}

type GetWorkspacesWorkspaceIdTagsRequest struct {
    WorkspaceID   string
}

type PatchCollectionsCollectionIdRequest struct {
    CollectionID   string
    Data     PatchCollectionsCollectionIDBody
}

type PatchScimV2GroupsGroupIdRequest struct {
    GroupID   string
    Data     PatchScimV2GroupsGroupIDBody
}

type PatchScimV2UsersUserIdRequest struct {
    UserID   string
    Data     PatchScimV2UsersUserIDBody
}

type PostApisRequest struct {
    WorkspaceID   string
    Data     PostApisBody
}

type PostApisApiIdCollectionsRequest struct {
    APIID   string
    Data     any
}

type PostApisApiIdSchemasRequest struct {
    APIID   string
    Data     PostApisAPIIDSchemasBody
}

type PostApisApiIdVersionsRequest struct {
    APIID   string
    Data     any
}

type PostCollectionsRequest struct {
    WorkspaceID   *string
    Data     PostCollectionsBody
}

type PostCollectionsForkCollectionIdRequest struct {
    CollectionID   string
    Workspace   string
    Data     PostCollectionsForkCollectionIDBody
}

type PostCollectionsMergeRequest struct {
    Data     PostCollectionsMergeBody
}

type PostCollectionsCollectionIdFoldersRequest struct {
    CollectionID   string
    Data     PostCollectionsCollectionIDFoldersBody
}

type PostCollectionsCollectionIdRequestsRequest struct {
    CollectionID   string
    FolderID   *string
    Data     PostCollectionsCollectionIDRequestsBody
}

type PostCollectionsCollectionIdResponsesRequest struct {
    CollectionID   string
    RequestID   string
    Data     PostCollectionsCollectionIDResponsesBody
}

type PostDetectedSecretsQueriesRequest struct {
    Cursor   *string
    Include   *string
    Limit   *int
    Data     PostDetectedSecretsQueriesBody
}

type PostEnvironmentsRequest struct {
    WorkspaceID   *string
    Data     PostEnvironmentsBody
}

type PostImportOpenapiRequest struct {
    WorkspaceID   *string
    Data     any
}

type PostMocksRequest struct {
    WorkspaceID   *string
    Data     PostMocksBody
}

type PostMocksMockIdPublishRequest struct {
    MockID   string
}

type PostMocksMockIdServerResponsesRequest struct {
    MockID   string
    Data     PostMocksMockIDServerResponsesBody
}

type PostMonitorsRequest struct {
    WorkspaceID   *string
    Data     PostMonitorsBody
}

type PostMonitorsMonitorIdRunRequest struct {
    MonitorID   string
}

type PostNetworkPrivateRequest struct {
    Data     any
}

type PostScimV2GroupsRequest struct {
    Data     PostScimV2GroupsBody
}

type PostScimV2UsersRequest struct {
    Data     PostScimV2UsersBody
}

type PostSecurityApiValidationRequest struct {
    Data     PostSecurityAPIValidationBody
}

type PostWebhooksRequest struct {
    WorkspaceID   *string
    Data     PostWebhooksBody
}

type PostWorkspacesRequest struct {
    Data     PostWorkspacesBody
}

type PutApisApiIdRequest struct {
    APIID   string
    Data     PutApisAPIIDBody
}

type PutApisApiIdCollectionsCollectionIdSyncWithSchemaTasksRequest struct {
    APIID   string
    CollectionID   string
}

type PutApisApiIdSchemasSchemaIdFilesFilePathRequest struct {
    APIID   string
    SchemaID   string
    FilePath   string
    Data     PutApisAPIIDSchemasSchemaIDFilesFilePathBody
}

type PutApisApiIdTagsRequest struct {
    APIID   string
    Data     PutApisAPIIDTagsBody
}

type PutApisApiIdVersionsVersionIdRequest struct {
    APIID   string
    VersionID   string
    Data     PutApisAPIIDVersionsVersionIDBody
}

type PutCollectionsCollectionIdRequest struct {
    CollectionID   string
    Data     PutCollectionsCollectionIDBody
}

type PutCollectionsCollectionIdFoldersFolderIdRequest struct {
    CollectionID   string
    FolderID   string
    Data     PutCollectionsCollectionIDFoldersFolderIDBody
}

type PutCollectionsCollectionIdRequestsRequestIdRequest struct {
    CollectionID   string
    RequestID   string
    Data     PutCollectionsCollectionIDRequestsRequestIDBody
}

type PutCollectionsCollectionIdResponsesResponseIdRequest struct {
    CollectionID   string
    ResponseID   string
    Data     PutCollectionsCollectionIDResponsesResponseIDBody
}

type PutCollectionsCollectionIdTagsRequest struct {
    CollectionID   string
    Data     PutCollectionsCollectionIDTagsBody
}

type PutDetectedSecretsSecretIdRequest struct {
    SecretID   string
    Data     PutDetectedSecretsSecretIDBody
}

type PutEnvironmentsEnvironmentIdRequest struct {
    EnvironmentID   string
    Data     PutEnvironmentsEnvironmentIDBody
}

type PutMocksMockIdRequest struct {
    MockID   string
    Data     PutMocksMockIDBody
}

type PutMocksMockIdServerResponsesServerResponseIdRequest struct {
    MockID   string
    ServerResponseID   string
    Data     PutMocksMockIDServerResponsesServerResponseIDBody
}

type PutMonitorsMonitorIdRequest struct {
    MonitorID   string
    Data     PutMonitorsMonitorIDBody
}

type PutNetworkPrivateNetworkEntityRequestRequestIdRequest struct {
    RequestID   int
    Data     PutNetworkPrivateNetworkEntityRequestRequestIDBody
}

type PutNetworkPrivateElementTypeElementIdRequest struct {
    ElementType   string
    ElementID   string
    Data     any
}

type PutScimV2UsersUserIdRequest struct {
    UserID   string
    Data     PutScimV2UsersUserIDBody
}

type PutWorkspacesWorkspaceIdRequest struct {
    WorkspaceID   string
    Data     PutWorkspacesWorkspaceIDBody
}

type PutWorkspacesWorkspaceIdGlobalVariablesRequest struct {
    WorkspaceID   string
    Data     PutWorkspacesWorkspaceIDGlobalVariablesBody
}

type PutWorkspacesWorkspaceIdTagsRequest struct {
    WorkspaceID   string
    Data     PutWorkspacesWorkspaceIDTagsBody
}
