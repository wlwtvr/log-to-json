import "@/styles/globals.css";

import { useEffect } from 'react';

function MyApp({ Component, pageProps }) {
  useEffect(() => {
    const loadWasm = async () => {
      const go = new Go();
      try {
        const wasm = await WebAssembly.instantiateStreaming(fetch('/wasm/main.wasm'), go.importObject);
        go.run(wasm.instance);
      } catch (err) {
        console.error('Failed to load WebAssembly module:', err);
      }
    };

    const script = document.createElement('script');
    script.src = '/wasm_exec.js';
    script.onload = loadWasm;
    document.body.appendChild(script);
  }, []);

  return <Component {...pageProps} />;
}

export default MyApp;
