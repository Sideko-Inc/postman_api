# postman_api go 

 The Postman API enables you to programmatically access data stored in your Postman account.

For a comprehensive set of examples of requests and responses, see the [**Postman API** collection](https://www.postman.com/postman/workspace/postman-public-workspace/documentation/12959542-c8142d51-e97c-46b6-bd77-52bb66712c9a).

## Important

- You must pass an `Accept` header with the `application/vnd.api.v10+json` value to use v10 and higher endpoints. While some of these endpoints may appear the same as the deprecated Postman v9 endpoints, they will use the v10 behavior when you send this `Accept` header. For more information, see [About v9 and v10 APIs](https://learning.postman.com/docs/developer/postman-api/intro-api/#about-v9-and-v10-apis).
- To use the **API** endpoints, you must first [update your APIs to the v10 format](https://learning.postman.com/docs/designing-and-developing-your-api/creating-an-api/#upgrading-an-api).

## Getting started

You can get started with the Postman API by [forking the Postman API collection](https://learning.postman.com/docs/collaborating-in-postman/version-control/#creating-a-fork) to your workspace. You can then use Postman to send requests.

## About the Postman API

- You must use a valid API Key to send requests to the API endpoints.
- The API has [rate and usage limits](https://learning.postman.com/docs/developer/postman-api/postman-api-rate-limits/).
- The API only responds to HTTPS-secured communications. Any requests sent via HTTP return an HTTP `301` redirect to the corresponding HTTPS resources.
- The API returns requests responses in [JSON format](https://en.wikipedia.org/wiki/JSON). When an API request returns an error, it is sent in the JSON response as an error key.
- The request method (verb) determines the nature of action you intend to perform. A request made using the `GET` method implies that you want to fetch something from Postman. The `POST` method implies you want to save something new to Postman.
- For all requests, API calls respond with their corresponding [HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). In the Postman client, the status code also provides help text that details the possible meaning of the response code.

### IDs and UIDs

All items in Postman, such as collections, workspaces, and APIs, have IDs and UIDs:

- An ID is the unique ID assigned to a Postman item. For example, `ec29121c-5203-409f-9e84-e83ffc10f226`.
- The UID is the **full** ID of a Postman item. This value is the item's unique ID concatenated with the user ID. For example, in the `12345678-ec29121c-5203-409f-9e84-e83ffc10f226` UID:
    - `12345678` is the user's ID.
    - `ec29121c-5203-409f-9e84-e83ffc10f226` is the item's ID.

### 503 response

An HTTP `503 Service Unavailable` response from our servers indicates there is an unexpected spike in API access traffic. The server is usually operational within the next five minutes. If the outage persists or you receive any other form of an HTTP `5XX` error, [contact support](https://support.postman.com/hc/en-us/requests/new/).

## Authentication

Postman uses API keys for authentication. The API key tells the API server that the request came from you. Everything that you have access to in Postman is accessible with your API key. You can [generate](https://learning.postman.com/docs/developer/postman-api/authentication/#generate-a-postman-api-key) a Postman API key in the [**API keys**](https://postman.postman.co/settings/me/api-keys) section of your Postman account settings.

You must include an API key in each request to the Postman API with the `X-Api-Key` request header. In Postman, you can store your API key as an [environment variable](https://www.getpostman.com/docs/environments). The Postman API [collection](https://www.getpostman.com/docs/collections) will use it to make API calls.

### Authentication error response

If an API key is missing, malformed, or invalid, you will receive an HTTP `401 Unauthorized` response code.

### Using the API key as a query parameter

Requests that accept the `X-Api-Key` request header also accept the API key when you send it as the `apikey` query parameter. An API key sent as part of the header has a higher priority when you send the key as both a request header and a query parameter.

## Rate and usage limits

API access [rate limits](https://learning.postman.com/docs/developer/postman-api/postman-api-rate-limits/) apply at a per-API key basis in unit time. The limit is **300 requests per minute**. Also, depending on your [plan](https://www.postman.com/pricing/), you may have usage limits. If you exceed either limit, your request will return an HTTP `429 Too Many Requests` status code.

Each API response returns the following set of headers to help you identify your use status:

| Header | Description |
| ------ | ----------- |
| `X-RateLimit-Limit` | The maximum number of requests that the consumer is permitted to make per minute. |
| `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window. |
| `X-RateLimit-Reset` | The time at which the current rate limit window resets in UTC epoch seconds. |

## Support

For help regarding accessing the Postman API, you can:

- Visit [Postman Support](https://support.postman.com/hc/en-us) or our [Community and Support](https://www.postman.com/community/) sites.
- Reach out to the [Postman community](https://community.postman.com/).
- Submit a help request to [Postman support](https://support.postman.com/hc/en-us/requests/new/).

## Policies

- [Postman Terms of Service](http://www.postman.com/legal/terms/)
- [Postman Privacy Policy](https://www.postman.com/legal/privacy-policy/)
 

 # Authentication 
 Get an API key from [postman](https://web.postman.co/)


And pass it to your SDK
 
 ```go
import "github.com/Sideko-Inc/postman_api"
client := postman_api.NewSidekoClient("API_KEY")
```

# delete_api
Deletes an API.
```go
response, err := client.DeleteApi(postman_api.DeleteApisApiIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# delete_schema_file
Deletes a file in an API schema.
```go
response, err := client.DeleteSchemaFile(postman_api.DeleteApisApiIdSchemasSchemaIdFilesFilePathRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    SchemaID: "5381f010-c4c1-11ed-afa1-0242ac120002",
    FilePath: "postman/collection/c1.json",
})
```
# delete_api_version
Deletes an API version.

**Note:**

This endpoint returns an HTTP `404 Not Found` response when an API version is pending publication.

```go
response, err := client.DeleteApiVersion(postman_api.DeleteApisApiIdVersionsVersionIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    VersionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# delete_collection
Deletes a collection.
```go
response, err := client.DeleteCollection(postman_api.DeleteCollectionsCollectionIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# delete_collection_folder
Deletes a folder in a collection.
```go
response, err := client.DeleteCollectionFolder(postman_api.DeleteCollectionsCollectionIdFoldersFolderIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    FolderID: "65a99e60-8e0a-4b6e-b79c-7d8264cc5caa",
})
```
# delete_collection_request
Deletes a request in a collection.
```go
response, err := client.DeleteCollectionRequest(postman_api.DeleteCollectionsCollectionIdRequestsRequestIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    RequestID: "c82dd02c-4870-4907-8fcb-593a876cf05b",
})
```
# delete_collection_response
Deletes a response in a collection.
```go
response, err := client.DeleteCollectionResponse(postman_api.DeleteCollectionsCollectionIdResponsesResponseIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    ResponseID: "cc364734-7dfd-4bfc-897d-be763dcdbb07",
})
```
# delete_environment
Deletes an environment.
```go
response, err := client.DeleteEnvironment(postman_api.DeleteEnvironmentsEnvironmentIdRequest{
    EnvironmentID: "5daabc50-8451-43f6-922d-96b403b4f28e",
})
```
# delete_mock
Deletes a mock server.
```go
response, err := client.DeleteMock(postman_api.DeleteMocksMockIdRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# delete_mock_server_response
Deletes a mock server's server response.
```go
response, err := client.DeleteMockServerResponse(postman_api.DeleteMocksMockIdServerResponsesServerResponseIdRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
    ServerResponseID: "965cdd16-fe22-4d96-a161-3d05490ac421",
})
```
# unpublish_mock
Unpublishes a mock server. Unpublishing a mock server sets its **Access Control** configuration setting to private.
```go
response, err := client.UnpublishMock(postman_api.DeleteMocksMockIdUnpublishRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# delete_monitor
Deletes a monitor.
```go
response, err := client.DeleteMonitor(postman_api.DeleteMonitorsMonitorIdRequest{
    MonitorID: "1e6b6cc1-c760-48e0-968f-4bfaeeae9af1",
})
```
# remove_element_or_folder
Removes an element or delete a folder from your [Private API Network](https://learning.postman.com/docs/collaborating-in-postman/adding-private-network/).

**Note:**

Removing an API, collection, or workspace element does **not** delete it. It only removes it from the Private API Network folder.
```go
response, err := client.RemoveElementOrFolder(postman_api.DeleteNetworkPrivateElementTypeElementIdRequest{
    ElementType: "api",
    ElementID: "5360b75f-447e-467c-9299-12fd6c92450d",
})
```
# delete_group
Deletes a group in Postman.

User accounts that were in the deleted group are deactivated in Postman if the app is assigned to the user only with the deleted group.

User accounts and the data corresponding to them are **not** deleted. To permanently delete user accounts and their data, [contact Postman support](https://www.postman.com/support/).

```go
response, err := client.DeleteGroup(postman_api.DeleteScimV2GroupsGroupIdRequest{
    GroupID: "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99",
})
```
# delete_workspace
Deletes an existing workspace.

### Important

If you delete a workspace that has a linked collection or environment with another workspace, this will delete the collection and environment in **all** workspaces.

```go
response, err := client.DeleteWorkspace(postman_api.DeleteWorkspacesWorkspaceIdRequest{
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# get_all_apis
Gets information about all APIs in a workspace.

**Note:**

This endpoint only returns APIs created or migrated in Postman v10 and higher.

```go
response, err := client.GetAllApis(postman_api.GetApisRequest{
    WorkspaceID: "9a7bb368-c4c4-11ed-afa1-0242ac120002",
})
```
# get_an_api
Gets information about an API.

**Note:**

- Git-connected APIs will **only** return the `versions` and `gitInfo` query responses. This is because schema and collection information is stored in the connected Git repository. The `gitInfo` object only lists the repository and folder locations of the files.
- API viewers can only use the `versions` option in the `include` query parameter.

```go
response, err := client.GetAnApi(postman_api.GetApisApiIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# get_collection
Gets a collection attached to an API. You can use the `versionId` query parameter to get a collection published in a version.

**Note:**

The `versionId` query parameter is a required parameter for API viewers.

```go
response, err := client.GetCollection(postman_api.GetApisApiIdCollectionsCollectionIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    CollectionID: "12345678-61867bcc-c4c1-11ed-afa1-0242ac120002",
})
```
# get_schema
Gets information about API schema. You can use the `versionId` query parameter to get a schema published in an API version.

You can use this API to do the following:

- Get a schema's metadata.
- Get all the files in a schema. This only returns the first file in the schema. The endpoint response contains a link to the next set of response results.
- Get a schema's contents in multi-file or bundled format.

**Note:**

The `versionId` query parameter is a **required** parameter for API viewers.

```go
response, err := client.GetSchema(postman_api.GetApisApiIdSchemasSchemaIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    SchemaID: "5381f010-c4c1-11ed-afa1-0242ac120002",
})
```
# get_schema_files
Gets the files in an API schema. You can use the `versionId` query parameter to get schema files published in an API version.

**Note:**

The `versionId` query parameter is a required parameter for API viewers.

```go
response, err := client.GetSchemaFiles(postman_api.GetApisApiIdSchemasSchemaIdFilesRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    SchemaID: "5381f010-c4c1-11ed-afa1-0242ac120002",
})
```
# get_schema_file_contents
Gets an API schema file contents at the defined path. You can use the `versionId` query parameter to get schema file contents published in an API version.

**Note:**

The `versionId` query parameter is a required parameter for API viewers.

```go
response, err := client.GetSchemaFileContents(postman_api.GetApisApiIdSchemasSchemaIdFilesFilePathRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    SchemaID: "5381f010-c4c1-11ed-afa1-0242ac120002",
    FilePath: "postman/collection/c1.json",
})
```
# get_api_tags
Gets all the tags associated with an API.
```go
response, err := client.GetApiTags(postman_api.GetApisApiIdTagsRequest{
    APIID: "12345678-6fd634a3-79ba-451d-8f07-56a953f96667",
})
```
# get_status_of_an_async_task
Gets the status of an asynchronous task.
```go
response, err := client.GetStatusOfAnAsyncTask(postman_api.GetApisApiIdTasksTaskIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    TaskID: "90ca9f5a-c4c4-21ed-afa1-0242ac120002",
})
```
# get_all_versions
Gets all the published versions of an API.
```go
response, err := client.GetAllVersions(postman_api.GetApisApiIdVersionsRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# get_api_version
Gets information about an API version.

**Note:**

- For API editors, this endpoint returns an HTTP `302 Found` status code when the version status is pending. It also returns the `/apis/{apiId}/tasks/{taskId}` task status response header.
- For API viewers, this endpoint returns an HTTP `404 Not Found` when the version status is pending.

```go
response, err := client.GetApiVersion(postman_api.GetApisApiIdVersionsVersionIdRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    VersionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# get_audit_logs
Gets a list of your team's generated audit events. For a complete list of all audit events, read our [Utilizing audit logs](https://learning.postman.com/docs/administration/audit-logs/) documentation.
```go
response, err := client.GetAuditLogs(postman_api.GetAuditLogsRequest{
})
```
# all_collections
Gets all of your [collections](https://www.getpostman.com/docs/collections). The response includes all of your subscribed collections.
```go
response, err := client.AllCollections(postman_api.GetCollectionsRequest{
})
```
# single_collection
Gets information about a collection. For a complete list of this endpoint's possible values, use the [collection.json schema file](https://schema.postman.com/json/collection/v2.1.0/collection.json).
```go
response, err := client.SingleCollection(postman_api.GetCollectionsCollectionIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# get_collection_folder
Gets information about a folder in a collection.
```go
response, err := client.GetCollectionFolder(postman_api.GetCollectionsCollectionIdFoldersFolderIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    FolderID: "65a99e60-8e0a-4b6e-b79c-7d8264cc5caa",
})
```
# get_collection_request
Gets information about a request in a collection.
```go
response, err := client.GetCollectionRequest(postman_api.GetCollectionsCollectionIdRequestsRequestIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    RequestID: "c82dd02c-4870-4907-8fcb-593a876cf05b",
})
```
# get_collection_response
Gets information about a response in a collection.
```go
response, err := client.GetCollectionResponse(postman_api.GetCollectionsCollectionIdResponsesResponseIdRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    ResponseID: "cc364734-7dfd-4bfc-897d-be763dcdbb07",
})
```
# get_collection_tags
Gets all the tags associated with a collection.
```go
response, err := client.GetCollectionTags(postman_api.GetCollectionsCollectionIdTagsRequest{
    CollectionID: "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# transform_collection_to_open_api
Transforms an existing Postman Collection into a stringified OpenAPI definition.

**Note:**

This does **not** create an API.

```go
response, err := client.TransformCollectionToOpenApi(postman_api.GetCollectionsCollectionIdTransformationsRequest{
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# get_detected_secrets_locations
Gets the locations of secrets detected by Postman's [Secret Scanner](https://learning.postman.com/docs/administration/secret-scanner/).
```go
response, err := client.GetDetectedSecretsLocations(postman_api.GetDetectedSecretsSecretIdLocationsRequest{
    SecretID: "ODk0MTU2",
    WorkspaceID: "e361eeb4-00dd-4225-9774-6146a2555999",
})
```
# all_environments
Gets information about all of your [environments](https://learning.postman.com/docs/sending-requests/managing-environments/).
```go
response, err := client.AllEnvironments(postman_api.GetEnvironmentsRequest{
})
```
# single_environment
Gets information about an environment.
```go
response, err := client.SingleEnvironment(postman_api.GetEnvironmentsEnvironmentIdRequest{
    EnvironmentID: "5daabc50-8451-43f6-922d-96b403b4f28e",
})
```
# api_key_owner
Gets information about the authenticated user.

**Note:**

This API returns a different response for users with the [Guest role](https://learning.postman.com/docs/collaborating-in-postman/roles-and-permissions/#team-roles).

```go
response, err := client.ApiKeyOwner()
```
# get_mocks
Gets all mock servers. By default, this endpoint returns only mock servers you created across all workspaces.

**Note:**

If you pass both the `teamId` and `workspace` query parameters, this endpoint only accepts the `workspace` query.

```go
response, err := client.GetMocks(postman_api.GetMocksRequest{
})
```
# get_mock
Gets information about a mock server.
```go
response, err := client.GetMock(postman_api.GetMocksMockIdRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# get_mock_call_logs
Gets a mock server's call logs. You can get a maximum of 6.5MB of call logs or a total of 100 call logs, whichever limit is met first in one API call.

Call logs contain exchanged request and response data made to mock servers. The logs provide visibility into how the mock servers are being used. You can log data to debug, test, analyze, and more, depending upon the use case.

```go
response, err := client.GetMockCallLogs(postman_api.GetMocksMockIdCallLogsRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# get_mock_server_responses
Gets all of a mock server's server responses.
```go
response, err := client.GetMockServerResponses(postman_api.GetMocksMockIdServerResponsesRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# get_mock_server_response
Gets information about a server response.
```go
response, err := client.GetMockServerResponse(postman_api.GetMocksMockIdServerResponsesServerResponseIdRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
    ServerResponseID: "965cdd16-fe22-4d96-a161-3d05490ac421",
})
```
# all_monitors
Gets all monitors.
```go
response, err := client.AllMonitors(postman_api.GetMonitorsRequest{
})
```
# single_monitor
Gets information about a monitor.
```go
response, err := client.SingleMonitor(postman_api.GetMonitorsMonitorIdRequest{
    MonitorID: "1e6b6cc1-c760-48e0-968f-4bfaeeae9af1",
})
```
# get_all_elements_and_folders
Gets information about the folders and their elements added to your [Private API Network](https://learning.postman.com/docs/collaborating-in-postman/adding-private-network/).

**Note:**

The `limit` and `offset` parameters are separately applied to elements and folders. For example, if you query a `limit` value of `10` and an `offset` value `0`, the endpoint returns 10 elements and 10 folders for a total of 20 items. The `totalCount` property in the `meta` response is the total count of **both** elements and folders.
```go
response, err := client.GetAllElementsAndFolders(postman_api.GetNetworkPrivateRequest{
})
```
# get_all_add_element_requests
Gets a list requests to add elements to the [Private API Network](https://learning.postman.com/docs/collaborating-in-postman/adding-private-network/).
```go
response, err := client.GetAllAddElementRequests(postman_api.GetNetworkPrivateNetworkEntityRequestAllRequest{
})
```
# fetch_all_group_resources
Gets information about all Postman team members.
```go
response, err := client.FetchAllGroupResources(postman_api.GetScimV2GroupsRequest{
})
```
# fetch_group_resource
Gets information about a Postman group within the team.
```go
response, err := client.FetchGroupResource(postman_api.GetScimV2GroupsGroupIdRequest{
    GroupID: "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99",
})
```
# get_resource_types
Gets all the resource types supported by Postman's SCIM API.
```go
response, err := client.GetResourceTypes()
```
# service_provider_config
Gets the Postman SCIM API configuration information. This includes a list of supported operations.
```go
response, err := client.ServiceProviderConfig()
```
# fetch_all_user_resources
Gets information about all Postman team members.
```go
response, err := client.FetchAllUserResources(postman_api.GetScimV2UsersRequest{
})
```
# fetch_user_resource
Gets information about a Postman team member.
```go
response, err := client.FetchUserResource(postman_api.GetScimV2UsersUserIdRequest{
    UserID: "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99",
})
```
# get_secret_types
Gets the metadata of the secret types supported by Postman's [Secret Scanner](https://learning.postman.com/docs/administration/secret-scanner/). You can use a secret type's ID in the response to query data with the POST `/detected-secrets/{secretId}` endpoint.
```go
response, err := client.GetSecretTypes()
```
# get_tagged_entities
Gets Postman elements (entities) by a given tag. Tags enable you to organize and search [workspaces](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/managing-workspaces/#tagging-a-workspace), [APIs](https://learning.postman.com/docs/designing-and-developing-your-api/managing-apis/#tagging-apis), and [collections](https://learning.postman.com/docs/collections/using-collections/#tagging-a-collection) that contain shared tags.

**Note:**

Tagging is available on [Postman Enterprise plans](https://www.postman.com/pricing/).
```go
response, err := client.GetTaggedEntities(postman_api.GetTagsSlugEntitiesRequest{
    Slug: "needs-review",
})
```
# all_workspaces
Gets all [workspaces](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/creating-workspaces/). The response includes your workspaces and any workspaces that you have access to.

**Note:**

This endpoint's response contains the visibility field. Visibility determines who can access the workspace:

- `personal` — Only you can access the workspace.
- `team` — All team members can access the workspace.
- `private` — Only invited team members can access the workspace ([Professional and Enterprise plans only](https://www.postman.com/pricing)).
- `public` — Everyone can access the workspace.
- `partner` — Only invited team members and [partners](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/partner-workspaces/) can access the workspace ([Enterprise Ultimate plans](https://www.postman.com/pricing) only).

```go
response, err := client.AllWorkspaces(postman_api.GetWorkspacesRequest{
})
```
# single_workspace
Gets information about a workspace.

**Note:**

This endpoint's response contains the `visibility` field. [Visibility](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/managing-workspaces/#changing-workspace-visibility) determines who can access the workspace:

- `personal` — Only you can access the workspace.
- `team` — All team members can access the workspace.
- `private` — Only invited team members can access the workspace ([Professional and Enterprise plans only](https://www.postman.com/pricing)).
- `public` — Everyone can access the workspace.
- `partner` — Only invited team members and [partners](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/partner-workspaces/) can access the workspace ([Enterprise Ultimate plans](https://www.postman.com/pricing) only).

### Important

We have **deprecated** the `name` and `uid` responses in the following array of objects:

- `collections`
- `environments`
- `mocks`
- `monitors`
- `apis`

```go
response, err := client.SingleWorkspace(postman_api.GetWorkspacesWorkspaceIdRequest{
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# get_workspace_global_variables
Gets a workspace's global [variables](https://learning.postman.com/docs/sending-requests/variables/#variable-scopes).
```go
response, err := client.GetWorkspaceGlobalVariables(postman_api.GetWorkspacesWorkspaceIdGlobalVariablesRequest{
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# get_workspace_tags
Gets all the tags associated with a workspace.
```go
response, err := client.GetWorkspaceTags(postman_api.GetWorkspacesWorkspaceIdTagsRequest{
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# patch_collection
Updates specific collection information, such as its name, events, or its variables. For more information about the `auth`, `variables`, and `events` properties, refer to the [collection.json schema file](https://schema.postman.com/json/collection/v2.1.0/collection.json):

- For `variables`, refer to `"#/definitions/variable"`.
- For `auth`, refer to `"#/definitions/auth-attribute"`.
- For `events`, refer to `"#/definitions/event"`.

For more information about the Collection Format, see the [Postman Collection Format documentation](https://learning.postman.com/collection-format/getting-started/overview/).

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPatchCollectionsCollectionIDBody(jsonBytes)

response, err := client.PatchCollection(postman_api.PatchCollectionsCollectionIdRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# patch_scim_v2_groups_group_id
Updates a group's information. Using this endpoint you can:

- Update a group's name.
- Add or remove members from a Postman group.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPatchScimV2GroupsGroupIDBody(jsonBytes)

response, err := client.PatchScimV2GroupsGroupId(postman_api.PatchScimV2GroupsGroupIdRequest{
    Data: data,
    GroupID: "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99",
})
```
# update_user_state
Updates a user's active state in Postman.

### Reactivating users

By setting the `active` property from `false` to `true`, this reactivates an account. This allows the account to authenticate in to Postman and adds the account back on to your Postman team.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPatchScimV2UsersUserIDBody(jsonBytes)

response, err := client.UpdateUserState(postman_api.PatchScimV2UsersUserIdRequest{
    Data: data,
    UserID: "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99",
})
```
# create_api
Creates an API.
```go
jsonBytes := []byte(`{
  "name": "Test API"
}`)
data, _ := postman_api.UnmarshalPostApisBody(jsonBytes)

response, err := client.CreateApi(postman_api.PostApisRequest{
    Data: data,
    WorkspaceID: "9a7bb368-c4c4-11ed-afa1-0242ac120002",
})
```
# add_collection
Adds a collection to an API. To do this, use the following `operationType` values:

- `COPY_COLLECTION` — Copies a collection from the workspace and adds it to an API.
- `CREATE_NEW` — Creates a new collection by providing the new collection's content.
- `GENERATE_FROM_SCHEMA` — Generates the collection from an API schema.
    - `options` — An object that contains advanced creation options and their values. You can find a complete list of properties and their values in Postman's OpenAPI 3.0 to Postman Collection v2.1.0 Converter OPTIONS documentation. These properties are case-sensitive.

```go
dataBytes := []byte(`{}`)
var data any
_ = json.Unmarshal(dataBytes, &data)

response, err := client.AddCollection(postman_api.PostApisApiIdCollectionsRequest{
    Data: data,
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# create_api_schema
Creates a schema for an API.
```go
jsonBytes := []byte(`{
  "files": [
    {}
  ],
  "type": "openapi:3"
}`)
data, _ := postman_api.UnmarshalPostApisAPIIDSchemasBody(jsonBytes)

response, err := client.CreateApiSchema(postman_api.PostApisApiIdSchemasRequest{
    Data: data,
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# create_api_version
Creates a new API version asynchronously and immediately returns an HTTP `202 Accepted` response. The response contains a polling link to the task status API in the `Location` header.

This endpoint is equivalent to publishing a version in Postman app, which is the snapshot of API collections and schema at a given point in time.

```go
dataBytes := []byte(`{
  "collections": [
    {}
  ],
  "name": "v1",
  "schemas": [
    {}
  ]
}`)
var data any
_ = json.Unmarshal(dataBytes, &data)

response, err := client.CreateApiVersion(postman_api.PostApisApiIdVersionsRequest{
    Data: data,
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# create_collection
Creates a collection using the [Postman Collection v2 schema format](https://schema.postman.com/json/collection/v2.1.0/docs/index.html).

For more information about the Collection Format, see the [Postman Collection Format documentation](https://learning.postman.com/collection-format/getting-started/overview/).

**Note:**

- For a complete list of available property values for this endpoint, use the following references available in the [collection.json schema file](https://schema.postman.com/json/collection/v2.1.0/collection.json):
    - `info` object — Use the `definitions.info` entry.
    - `item` object — Use the `definitions.items` entry.
- For all other possible values, refer to the [collection.json schema file](https://schema.postman.com/json/collection/v2.1.0/collection.json).

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostCollectionsBody(jsonBytes)

response, err := client.CreateCollection(postman_api.PostCollectionsRequest{
    Data: data,
})
```
# create_a_fork
Creates a [fork](https://learning.postman.com/docs/collaborating-in-postman/version-control/#creating-a-fork) from an existing collection into a workspace.
```go
jsonBytes := []byte(`{
  "label": "Test Fork"
}`)
data, _ := postman_api.UnmarshalPostCollectionsForkCollectionIDBody(jsonBytes)

response, err := client.CreateAFork(postman_api.PostCollectionsForkCollectionIdRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    Workspace: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# merge_a_fork
Merges a forked collection back into its destination collection.
```go
jsonBytes := []byte(`{
  "destination": "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2",
  "source": "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2"
}`)
data, _ := postman_api.UnmarshalPostCollectionsMergeBody(jsonBytes)

response, err := client.MergeAFork(postman_api.PostCollectionsMergeRequest{
    Data: data,
})
```
# create_collection_folder
Creates a folder in a collection. For a complete list of properties, refer to "Folder" in the [collection.json schema file](https://schema.postman.com/collection/json/v2.1.0/draft-07/collection.json).

You can use this endpoint to to import requests and responses into a newly-created folder. To do this, include the `requests` field and the list of request objects in the request body. For more information, see the provided example.

**Note:**

It is recommended that you pass the `name` property in the request body. If you do not, the system uses a null value. As a result, this creates a folder with a blank name.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostCollectionsCollectionIDFoldersBody(jsonBytes)

response, err := client.CreateCollectionFolder(postman_api.PostCollectionsCollectionIdFoldersRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# create_collection_request
Creates a request in a collection. For a complete list of properties, see the [Collection Format Request documentation](https://learning.postman.com/collection-format/reference/request/).

**Note:**

It is recommended that you pass the `name` property in the request body. If you do not, the system uses a null value. As a result, this creates a request with a blank name.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostCollectionsCollectionIDRequestsBody(jsonBytes)

response, err := client.CreateCollectionRequest(postman_api.PostCollectionsCollectionIdRequestsRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# create_collection_response
Creates a request response in a collection. For a complete list of properties, see the [Collection Format Response documentation](https://learning.postman.com/collection-format/reference/response/#reference-diagram).

**Note:**

It is recommended that you pass the `name` property in the request body. If you do not, the system uses a null value. As a result, this creates a response with a blank name.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostCollectionsCollectionIDResponsesBody(jsonBytes)

response, err := client.CreateCollectionResponse(postman_api.PostCollectionsCollectionIdResponsesRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    RequestID: "string",
})
```
# detected_secrets_queries
Returns all secrets detected by Postman's [Secret Scanner](https://learning.postman.com/docs/administration/secret-scanner/), grouped by workspace. If you pass an empty request body, this endpoint returns all results.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostDetectedSecretsQueriesBody(jsonBytes)

response, err := client.DetectedSecretsQueries(postman_api.PostDetectedSecretsQueriesRequest{
    Data: data,
})
```
# create_environment
Creates an environment.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostEnvironmentsBody(jsonBytes)

response, err := client.CreateEnvironment(postman_api.PostEnvironmentsRequest{
    Data: data,
})
```
# import_external_api_specification
Imports an OpenAPI definition into Postman as a new [Postman Collection](https://learning.postman.com/docs/getting-started/creating-the-first-collection/).
```go
dataBytes := []byte(`{}`)
var data any
_ = json.Unmarshal(dataBytes, &data)

response, err := client.ImportExternalApiSpecification(postman_api.PostImportOpenapiRequest{
    Data: data,
})
```
# create_mock
**In Postman v10 and higher you cannot create mocks for collections added to an API definition.**

Creates a mock server in a collection.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostMocksBody(jsonBytes)

response, err := client.CreateMock(postman_api.PostMocksRequest{
    Data: data,
})
```
# publish_mock
Publishes a mock server. Publishing a mock server sets its **Access Control** configuration setting to public.
```go
response, err := client.PublishMock(postman_api.PostMocksMockIdPublishRequest{
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# create_server_response
Creates a server response. Server responses let you simulate 5xx server-level responses, such as 500 or 503.

Server-level responses are agnostic to application-level logic. Server responses let you simulate this behavior on a mock server. You do not need to define each error for all exposed paths on the mock server.

If you set a server response as active, then all the calls to the mock server return with that active server response.

**Note:**

You can create multiple server responses for a mock server, but only one mock server can be set as active.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostMocksMockIDServerResponsesBody(jsonBytes)

response, err := client.CreateServerResponse(postman_api.PostMocksMockIdServerResponsesRequest{
    Data: data,
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# create_monitor
**In Postman v10 and higher you cannot create monitors for collections added to an API definition.**

Creates a monitor.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostMonitorsBody(jsonBytes)

response, err := client.CreateMonitor(postman_api.PostMonitorsRequest{
    Data: data,
})
```
# run_a_monitor
Runs a monitor and returns its run results.
```go
response, err := client.RunAMonitor(postman_api.PostMonitorsMonitorIdRunRequest{
    MonitorID: "1e6b6cc1-c760-48e0-968f-4bfaeeae9af1",
})
```
# post_element_or_folder
Publishes a element or creates a folder in your [Private API Network](https://learning.postman.com/docs/collaborating-in-postman/adding-private-network/). An element is a Postman API, collection, or workspace.
```go
dataBytes := []byte(`{}`)
var data any
_ = json.Unmarshal(dataBytes, &data)

response, err := client.PostElementOrFolder(postman_api.PostNetworkPrivateRequest{
    Data: data,
})
```
# create_group
Creates a new user group in Postman and creates a new account for each group member.

Each account is added to your Postman team and authentication is activated for each user. If an existing Postman account uses an email that matches a group member's email ID, an [email invite](https://postman.postman.co/docs/administration/managing-your-team/managing-your-team/#invites) to join your Postman team is sent to that user. Once the user accepts the invite, they'll be added to your team.

By default, the system assigns new users the developer role. You can [update user roles in Postman](https://learning.postman.com/docs/administration/managing-your-team/managing-your-team/#managing-team-roles).

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostScimV2GroupsBody(jsonBytes)

response, err := client.CreateGroup(postman_api.PostScimV2GroupsRequest{
    Data: data,
})
```
# create_user
Creates a new user account in Postman and adds the user to your organization's Postman team. If the account does not already exist, this also activates the user so they can authenticate in to your Postman team.

If the account already exists, the system sends the user an [email invite](https://learning.postman.com/docs/administration/managing-your-team/managing-your-team/#inviting-users) to join the Postman team. The user joins the team once they accept the invite.

By default, the system assigns new users the developer role. You can [update user roles in Postman](https://learning.postman.com/docs/administration/managing-your-team/managing-your-team/#managing-team-roles).

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostScimV2UsersBody(jsonBytes)

response, err := client.CreateUser(postman_api.PostScimV2UsersRequest{
    Data: data,
})
```
# schema_security_validation
Performs an analysis on the given definition and returns any issues based on your [predefined rulesets](https://learning.postman.com/docs/api-governance/configurable-rules/configurable-rules-overview/). This endpoint can help you understand the violations' impact and offers solutions to help you resolve any errors. You can include this endpoint to your CI/CD process to automate schema validation.

For more information, see our [Rule violations in the API definition](https://learning.postman.com/docs/api-governance/api-definition/api-definition-warnings/) documentation.

**Note:**

- The maximum allowed size of the definition is 10 MB.
- You must [import and enable](https://learning.postman.com/docs/api-governance/configurable-rules/configuring-api-security-rules/) [Postman's OWASP security rules](https://postman.postman.co/api-governance/libraries/postman_owasp/view) in Postman for this endpoint to return any security rule violations.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostSecurityAPIValidationBody(jsonBytes)

response, err := client.SchemaSecurityValidation(postman_api.PostSecurityApiValidationRequest{
    Data: data,
})
```
# create_webhook
Creates a webhook that triggers a collection with a custom payload. You can get the webhook's URL from the `webhookUrl` property in the endpoint's response.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostWebhooksBody(jsonBytes)

response, err := client.CreateWebhook(postman_api.PostWebhooksRequest{
    Data: data,
})
```
# create_workspace
Creates a new [workspace](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/creating-workspaces/).

### Important

We **deprecated** linking collections or environments between workspaces. We do **not** recommend that you do this.

If you have a linked collection or environment, note the following:

- The endpoint does **not** create a clone of a collection or environment.
- Any changes you make to a linked collection or environment changes them in **all** workspaces.
- If you delete a collection or environment linked between workspaces, the system deletes it in **all** the workspaces.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPostWorkspacesBody(jsonBytes)

response, err := client.CreateWorkspace(postman_api.PostWorkspacesRequest{
    Data: data,
})
```
# update_an_api
Updates an API.
```go
jsonBytes := []byte(`{
  "name": "Test API"
}`)
data, _ := postman_api.UnmarshalPutApisAPIIDBody(jsonBytes)

response, err := client.UpdateAnApi(postman_api.PutApisApiIdRequest{
    Data: data,
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
})
```
# sync_collection_with_schema
Syncs a collection attached to an API with the API schema.

This is an asynchronous endpoint that returns an HTTP `202 Accepted` response. The response contains a polling link to the `/apis/{apiId}/tasks/{taskId}` endpoint in the `Location` header.

**Note:**

This endpoint only supports the OpenAPI 3 schema type.

```go
response, err := client.SyncCollectionWithSchema(postman_api.PutApisApiIdCollectionsCollectionIdSyncWithSchemaTasksRequest{
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    CollectionID: "12345678-61867bcc-c4c1-11ed-afa1-0242ac120002",
})
```
# create_or_update_schema_file
Creates or updates an API schema file.

**Note:**

- If the provided file path exists, the file will be updated with the new contents.
- If the provided file path does **not** exist, then a new schema file will be created.
- If the file path contains a `/` (forward slash) character, then a folder is created. For example, if the file path is the `dir/schema.json` value, then a `dir` folder is created with the `schema.json` file inside.

```go
jsonBytes := []byte(`{
  "content": "string"
}`)
data, _ := postman_api.UnmarshalPutApisAPIIDSchemasSchemaIDFilesFilePathBody(jsonBytes)

response, err := client.CreateOrUpdateSchemaFile(postman_api.PutApisApiIdSchemasSchemaIdFilesFilePathRequest{
    Data: data,
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    SchemaID: "5381f010-c4c1-11ed-afa1-0242ac120002",
    FilePath: "postman/collection/c1.json",
})
```
# update_api_tags
Updates an API's associated tags. This endpoint replaces all existing tags with those you pass in the request body.
```go
jsonBytes := []byte(`{
  "tags": [
    {
      "slug": "needs-review"
    }
  ]
}`)
data, _ := postman_api.UnmarshalPutApisAPIIDTagsBody(jsonBytes)

response, err := client.UpdateApiTags(postman_api.PutApisApiIdTagsRequest{
    Data: data,
    APIID: "12345678-6fd634a3-79ba-451d-8f07-56a953f96667",
})
```
# update_api_version
Updates an API version.

**Note:**

This endpoint returns an HTTP `404 Not Found` response when an API version is pending publication.

```go
jsonBytes := []byte(`{
  "name": "Release 1.5"
}`)
data, _ := postman_api.UnmarshalPutApisAPIIDVersionsVersionIDBody(jsonBytes)

response, err := client.UpdateApiVersion(postman_api.PutApisApiIdVersionsVersionIdRequest{
    Data: data,
    APIID: "90ca9f5a-c4c4-11ed-afa1-0242ac120002",
    VersionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# put_collection
Replaces the contents of a collection using the [Postman Collection v2 schema format](https://schema.postman.com/json/collection/v2.1.0/docs/index.html). Include the collection's ID values in the request body. If you do not, the endpoint removes the existing items and creates new items.

For a complete list of available property values for this endpoint, use the following references available in the [collection.json schema file](https://schema.postman.com/json/collection/v2.1.0/collection.json):
- `info` object — Use `"#/definitions/info"`.
- `item` object — Use `"#/definitions/item"`.

For all other possible values, refer to the [collection.json schema file](https://schema.postman.com/json/collection/v2.1.0/collection.json). For more information about the Collection Format, see the [Postman Collection Format documentation](https://learning.postman.com/collection-format/getting-started/overview/).

**Note**

To copy another collection's contents to the given collection, **remove** all ID values before you pass it in this endpoint. If you do not, this endpoint returns an error. These values include the `id`, `uid`, and `postman_id` values.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutCollectionsCollectionIDBody(jsonBytes)

response, err := client.PutCollection(postman_api.PutCollectionsCollectionIdRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# update_collection_folder
Updates a folder in a collection. For a complete list of properties, refer to "Folder" in the [collection.json schema file](https://schema.postman.com/collection/json/v2.1.0/draft-07/collection.json).

**Note:**

This endpoint acts like a PATCH method. It only updates the values that you pass in the request body (for example, the `name` property). The endpoint does **not** update the entire resource.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutCollectionsCollectionIDFoldersFolderIDBody(jsonBytes)

response, err := client.UpdateCollectionFolder(postman_api.PutCollectionsCollectionIdFoldersFolderIdRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    FolderID: "65a99e60-8e0a-4b6e-b79c-7d8264cc5caa",
})
```
# update_collection_request
Updates a request in a collection. For a complete list of properties, see the [Collection Format Request documentation](https://learning.postman.com/collection-format/reference/request/).

**Note:**

- You must pass a collection ID (`12ece9e1-2abf-4edc-8e34-de66e74114d2`), not a collection(`12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2`), in this endpoint.
- This endpoint does not support changing the folder of a request.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutCollectionsCollectionIDRequestsRequestIDBody(jsonBytes)

response, err := client.UpdateCollectionRequest(postman_api.PutCollectionsCollectionIdRequestsRequestIdRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    RequestID: "c82dd02c-4870-4907-8fcb-593a876cf05b",
})
```
# update_collection_response
Updates a response in a collection. For a complete list of properties, see the [Collection Format Response documentation](https://learning.postman.com/collection-format/reference/response/#reference-diagram).

**Note:**

- You must pass a collection ID (`12ece9e1-2abf-4edc-8e34-de66e74114d2`), not a collection UID (`12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2`), in this endpoint.
- This endpoint acts like a PATCH method. It only updates the values that you pass in the request body (for example, the `name` property). The endpoint does **not** update the entire resource.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutCollectionsCollectionIDResponsesResponseIDBody(jsonBytes)

response, err := client.UpdateCollectionResponse(postman_api.PutCollectionsCollectionIdResponsesResponseIdRequest{
    Data: data,
    CollectionID: "12ece9e1-2abf-4edc-8e34-de66e74114d2",
    ResponseID: "cc364734-7dfd-4bfc-897d-be763dcdbb07",
})
```
# update_collection_tags
Updates a collection's associated tags. This endpoint replaces all existing tags with those you pass in the request body.
```go
jsonBytes := []byte(`{
  "tags": [
    {
      "slug": "needs-review"
    }
  ]
}`)
data, _ := postman_api.UnmarshalPutCollectionsCollectionIDTagsBody(jsonBytes)

response, err := client.UpdateCollectionTags(postman_api.PutCollectionsCollectionIdTagsRequest{
    Data: data,
    CollectionID: "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2",
})
```
# update_detected_secret_resolutions
Updates the resolution status of a secret detected in a workspace.
```go
jsonBytes := []byte(`{
  "resolution": "ACCEPTED_RISK",
  "workspaceId": "e361eeb4-00dd-4225-9774-6146a2555999"
}`)
data, _ := postman_api.UnmarshalPutDetectedSecretsSecretIDBody(jsonBytes)

response, err := client.UpdateDetectedSecretResolutions(postman_api.PutDetectedSecretsSecretIdRequest{
    Data: data,
    SecretID: "ODk0MTU2",
})
```
# update_environment
Updates an environment.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutEnvironmentsEnvironmentIDBody(jsonBytes)

response, err := client.UpdateEnvironment(postman_api.PutEnvironmentsEnvironmentIdRequest{
    Data: data,
    EnvironmentID: "5daabc50-8451-43f6-922d-96b403b4f28e",
})
```
# update_mock
Updates a mock server.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutMocksMockIDBody(jsonBytes)

response, err := client.UpdateMock(postman_api.PutMocksMockIdRequest{
    Data: data,
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
})
```
# update_server_response
Updates a server response.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutMocksMockIDServerResponsesServerResponseIDBody(jsonBytes)

response, err := client.UpdateServerResponse(postman_api.PutMocksMockIdServerResponsesServerResponseIdRequest{
    Data: data,
    MockID: "e3d951bf-873f-49ac-a658-b2dcb91d3289",
    ServerResponseID: "965cdd16-fe22-4d96-a161-3d05490ac421",
})
```
# update_monitor
Updates a monitor.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutMonitorsMonitorIDBody(jsonBytes)

response, err := client.UpdateMonitor(postman_api.PutMonitorsMonitorIdRequest{
    Data: data,
    MonitorID: "1e6b6cc1-c760-48e0-968f-4bfaeeae9af1",
})
```
# respond_element_add_request
Responds to a request to add an element to the [Private API Network](https://learning.postman.com/docs/collaborating-in-postman/adding-private-network/). Only managers can approve or deny a request. Once approved, the element will appear in the team's Private API Network.
```go
jsonBytes := []byte(`{
  "status": "denied"
}`)
data, _ := postman_api.UnmarshalPutNetworkPrivateNetworkEntityRequestRequestIDBody(jsonBytes)

response, err := client.RespondElementAddRequest(postman_api.PutNetworkPrivateNetworkEntityRequestRequestIdRequest{
    Data: data,
    RequestID: 232,
})
```
# put_element_or_folder
Updates an element or folder in your [Private API Network](https://learning.postman.com/docs/collaborating-in-postman/adding-private-network/).
```go
dataBytes := []byte(`{}`)
var data any
_ = json.Unmarshal(dataBytes, &data)

response, err := client.PutElementOrFolder(postman_api.PutNetworkPrivateElementTypeElementIdRequest{
    Data: data,
    ElementType: "api",
    ElementID: "5360b75f-447e-467c-9299-12fd6c92450d",
})
```
# update_user_information
Updates a user's first and last name in Postman.

**Note:**

You can only use the SCIM API to update a user's first and last name. You cannot update any other user attributes with the API.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutScimV2UsersUserIDBody(jsonBytes)

response, err := client.UpdateUserInformation(postman_api.PutScimV2UsersUserIdRequest{
    Data: data,
    UserID: "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99",
})
```
# update_workspace
Updates a workspace.

### Important

We **deprecated** linking collections or environments between workspaces. We do **not** recommend that you do this.

If you have a linked collection or environment, note the following:

- The endpoint does **not** create a clone of a collection or environment.
- Any changes you make to a linked collection or environment changes them in **all** workspaces.
- If you delete a collection or environment linked between workspaces, the system deletes it in **all** the workspaces.

```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutWorkspacesWorkspaceIDBody(jsonBytes)

response, err := client.UpdateWorkspace(postman_api.PutWorkspacesWorkspaceIdRequest{
    Data: data,
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# put_workspace_global_variables
Updates and replaces a workspace's global [variables](https://learning.postman.com/docs/sending-requests/variables/#variable-scopes). This endpoint replaces all existing global variables with the variables you pass in the request body.
```go
jsonBytes := []byte(`{}`)
data, _ := postman_api.UnmarshalPutWorkspacesWorkspaceIDGlobalVariablesBody(jsonBytes)

response, err := client.PutWorkspaceGlobalVariables(postman_api.PutWorkspacesWorkspaceIdGlobalVariablesRequest{
    Data: data,
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
# update_workspace_tags
Updates a workspace's associated tags. This endpoint replaces all existing tags with those you pass in the request body.
```go
jsonBytes := []byte(`{
  "tags": [
    {
      "slug": "needs-review"
    }
  ]
}`)
data, _ := postman_api.UnmarshalPutWorkspacesWorkspaceIDTagsBody(jsonBytes)

response, err := client.UpdateWorkspaceTags(postman_api.PutWorkspacesWorkspaceIdTagsRequest{
    Data: data,
    WorkspaceID: "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
})
```
