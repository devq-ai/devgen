# Required FastMCP server pattern
from fastmcp import FastMCP
import logfire

# Configure MCP server with Logfire
@logfire.instrument()
async def create_mcp_server():
    mcp = FastMCP("server-name")
    
    @mcp.tool()
    async def example_tool(query: str) -> str:
        """Tool must return real results, never mock data"""
        with logfire.span('mcp-tool-execution'):
            # Real implementation required
            result = await actual_processing(query)
            return result
    
    return mcp

# Required MCP client pattern for agentical
from fastmcp.client import MCPClient

@logfire.instrument()
async def connect_mcp_client(server_url: str):
    client = MCPClient(server_url)
    # Test connection before returning
    await client.ping()
    return client
