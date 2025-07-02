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
			mcp.WithDescription("Searches Indian Kanoon for legal case documents (judgments and tribunal decisions) using a free-text query. Returns a list of matching documents, each with its unique document ID (TID), type (judgment or tribunal), publication date, document size, and headline. Useful for retrieving metadata about relevant legal cases based on keywords, party names, citations, or other search terms. The results can be used to fetch full document details with the document ID."),
			mcp.WithString("search_query",
				mcp.Description("A free-text search string to query Indian Kanoon. This can include keywords, party names, case numbers, citations, or any relevant legal search terms. The search is performed across judgments and tribunal decisions."),
				mcp.Required(),
			),
		),
		handleIndianKanoonSearch,
	)

	s.AddTool(
		mcp.NewTool("indian_kannon_fetch_document",
		mcp.WithDescription("Fetches the full details of a specific legal document from Indian Kanoon using its document ID (TID). Returns the document's title, publication date, full text, source information, document type, and whether it is an official court copy. Use this after obtaining a TID from a search to retrieve the complete case content and metadata."),
		mcp.WithString("document_id",
		mcp.Description("The unique document ID (TID) of the legal case to fetch. This ID is obtained from the search tool and is required to retrieve the full text and metadata of the selected case document."),
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
