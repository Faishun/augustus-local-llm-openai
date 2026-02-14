# Augustus - LLM Vulnerability Scanner

> Official repo for more commands/probes/docs: https://github.com/praetorian-inc/augustus

Run a simple quick test:
```bash
./bin/augustus scan openai.OpenAI \
  --probes-glob "dan.*,encoding.*" \
  --buff encoding.Base64 \
  --config '{"model":"qwen/qwen3-4b-2507","base_url":"http://localhost:8000/v1"}' \
  --verbose --format table
```
