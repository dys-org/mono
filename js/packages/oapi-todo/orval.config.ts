import { defineConfig } from 'orval';

export default defineConfig({
  api: {
    input: '../../../go/oapi-todo/api/openapi.yaml',
    output: {
      target: './src/api.ts',
      schemas: './src/model',
      client: 'vue-query',
      httpClient: 'fetch',
      baseUrl: '/api',
      prettier: true,
      // mock: true,
    },
  },
});
