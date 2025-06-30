package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lawAgent/indianKanoon"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func handleIndianKanoonSearch(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query := req.GetString("search_query", "")
	fmt.Println("received Request for query", query)
	ikApiClient := indianKanoon.GetIKApiClient()
	ikSearchData := indianKanoon.IKSearchData{
		FormInput: query,
		DocTypes:  indianKanoon.DOCTYPE,
	}

	response := ikApiClient.SearchQuery(ikSearchData)
	jsonBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal search response: %w", err)
	}
	return mcp.NewToolResultText(string(jsonBytes)), nil
}
func handleIndianKanoonDocumentFetch(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query := req.GetString("document_id", "")
	fmt.Println("received Request for query", query)
	ikApiClient := indianKanoon.GetIKApiClient()
	ikSearchDocument := indianKanoon.IKSearchDocument{
		DocId: query,
	}

	response := ikApiClient.DocumentFetch(ikSearchDocument)
	jsonBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal search response: %w", err)
	}
	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func main() {
	s := server.NewMCPServer(
		"indian Kannon MCP server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	s.AddTool(
		mcp.NewTool("indian_kannon_search",
			mcp.WithDescription("This is a tool to make search queries on Indian Kanoon which returns a list of TIDs (document IDs), doctype (judgments or tribunals), published date, docsize, and headline."),
			mcp.WithString("search_query",
				mcp.Description("The query string to search Indian Kanoon"),
				mcp.Required(),
			),
		),
		handleIndianKanoonSearch,
	)

	s.AddTool(
		mcp.NewTool("indian_kannon_fetch_document",
		mcp.WithDescription("Searches Indian Kanoon for legal documents and returns the document ID (TID), title, publication date, full text, source information, document type, and official court copy status."),
		mcp.WithString("document_ID",
		mcp.Description("The document ID (TID) to fetch the full case content"),
		mcp.Required(),
			),
		),
		handleIndianKanoonDocumentFetch,
	)

	fmt.Println("Starting HTTP server on :8080...")

	log.Println("Starting HTTP server on :8080")
	httpServer := server.NewStreamableHTTPServer(s)
	if err := httpServer.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
