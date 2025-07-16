import { GoogleGenAI, FunctionCallingConfigMode , mcpToTool} from '@google/genai';

const ai = new GoogleGenAI({
    apiKey:"AIzaSyAFq0CHvNCfHS2XBwtIcQFmtvnnlhldY4E"
});

export {ai} ;