import React, { useState } from 'react';
import ReactMarkdown from 'react-markdown';

const SearchBar = () => {
  const [searchInput, setSearchInput] = useState("");
  const [responseText, setResponseText] = useState("");

  const handleChange = (e) => {
    setSearchInput(e.target.value);
  };

  const handleSubmit = async () => {
    try {
      const res = await fetch("http://localhost:3000/api/generate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ prompt: searchInput }),
      });

      if (!res.ok) {
        throw new Error(`Server error: ${res.status}`);
      }

      const data = await res.json();
      console.log("Backend response:", data);
      setResponseText(data.response); // Show the response
    } catch (error) {
      console.error("Error calling /api/generate:", error);
      setResponseText("Error: " + error.message);
    }
  };

  return (
    <div>
      <input
        type="search"
        placeholder="Search legal case..."
        onChange={handleChange}
        value={searchInput}
      />
      <button onClick={handleSubmit}>Search</button>

      <div>
        <h3>Response:</h3>
        <ReactMarkdown>{responseText.replace(/\\n/g, '\n')}</ReactMarkdown>      </div>
    </div>
  );
};

export default SearchBar;
