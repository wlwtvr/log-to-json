import * as parserJson from 'prettier/parser-babel';
import * as prettierPluginEstree from "prettier/plugins/estree";
import * as prettier from 'prettier/standalone';
import { useState } from 'react';

const ParserComponent = () => {
  const [input, setInput] = useState('');
  const [output, setOutput] = useState('');
  const [error, setError] = useState('');

  const handleParse = async () => {
    setError('');
    try {
      await waitForWasm(); // Ensure the WASM module is loaded
      if (window.parseTextToJSON) {
        const result = window.parseTextToJSON(input);
        const prettyOutput = prettier.format(JSON.stringify(result, null, 2), {
          parser: 'json',
          plugins: [parserJson, prettierPluginEstree],
        });
        setOutput(prettyOutput);
      } else {
        setError('WebAssembly module is not loaded.');
      }
    } catch (err) {
      setError('Failed to parse input.');
    }
  };

  const waitForWasm = () => {
    return new Promise((resolve) => {
      const interval = setInterval(() => {
        if (window.parseTextToJSON) {
          clearInterval(interval);
          resolve();
        }
      }, 100);
    });
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(output);
    alert('Copied to clipboard');
  };

  return (
    <div className="p-4 max-w-4xl mx-auto">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <textarea
          className="p-2 border border-gray-300 rounded text-black"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Enter your text here"
          rows={20}
        />
        <div className="flex flex-col items-center justify-center">
          <button
            className="mb-2 bg-green-500 text-white px-4 py-2 rounded"
            onClick={handleParse}
          >
            Parse
          </button>
          <button
            className="bg-blue-500 text-white px-4 py-2 rounded"
            onClick={copyToClipboard}
          >
            Copy
          </button>
        </div>
        <textarea
          className="p-2 border border-gray-300 rounded text-black"
          value={output}
          readOnly
          placeholder="Parsed JSON will appear here"
          rows={20}
        />
      </div>
      {error && <p className="text-red-500 mt-2">{error}</p>}
    </div>
  );
};

export default ParserComponent;
