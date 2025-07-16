import { client, serverParams } from "./server/mcp_client.js";
import { mcpToTool } from "@google/genai";
import { ai } from "./server/gemini_client.js";
import express from "express";
import bodyParser from "body-parser";
import cors from "cors";
import path from "path";
import { fileURLToPath } from "url";

const app = express();
const PORT = 3000;

// Required for __dirname in ES Modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Middleware
app.use(bodyParser.json());
app.use(cors());

// Serve static files from frontend/dist
app.use(express.static(path.join(__dirname, 'frontend/dist')));

// Connect MCP client
await client.connect(serverParams);

// API route
app.post('/api/generate', async (req, res) => {
  try {
    const { prompt } = req.body;
    const tools = await client.listTools();
    console.log(tools);
    const response = await ai.models.generateContent({
      model: "gemini-2.5-flash",
      contents: "Please answer in markdown format." + prompt,
      config: {
        tools: [mcpToTool(client)],
      },
    });
    res.json({ response: response.text });
  } catch (error) {
    console.error(error);
    res.status(500).json({ error: error.message || "Internal Server Error" });
  }
});

 

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
