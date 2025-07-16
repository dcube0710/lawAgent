import { Client } from "@modelcontextprotocol/sdk/client/index.js";
import { StreamableHTTPClientTransport } from "@modelcontextprotocol/sdk/client/streamableHttp.js";

const BASE_URL = 'http://localhost:8081/mcp';

const serverParams = new StreamableHTTPClientTransport(BASE_URL);

const client = new Client({
    name:"MCP Client",
    version:"1.0.0"
});

export { serverParams, client };




