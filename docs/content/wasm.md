---
type: helper
layout: single
---
{{< rawhtml >}}
<!doctype html>
<html>

<head>
    <meta charset="utf-8">
</head>

<body style="overflow: hidden;">
    <!-- Polyfill for the old Edge browser -->
    <script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script>
    <script src="/js/wasm_exec.js"></script>
    <link rel="stylesheet" href="/css/iframe.css">
    <script>
        window.addEventListener('DOMContentLoaded', async () => {
            const go = new Go();
            const name = window.location.search.substring(1);
            let url = `./${name}.wasm`;

            // Polyfill
            if (!WebAssembly.instantiateStreaming) {
                WebAssembly.instantiateStreaming = async (resp, importObject) => {
                    const source = await (await resp).arrayBuffer();
                    return await WebAssembly.instantiate(source, importObject);
                };
            }

            const result = await WebAssembly.instantiateStreaming(await fetch(url), go.importObject).catch((err) => {
                console.error(err);
            });
            document.getElementById('loading').remove();
            go.run(result.instance);
        });
    </script>
    <p id="loading">Loading...</p>
</body>

</html>
{{< /rawhtml >}}