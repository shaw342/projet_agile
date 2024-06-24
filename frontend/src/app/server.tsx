const { createServer } = require('http');
const { parse } = require('url');
const next = require('next');
const proxy = require('http-proxy-middleware');

const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev });
const handle = app.getRequestHandler();

app.prepare().then(() => {
  createServer((req: { url: any; }, res: any) => {
    // Définissez le chemin d'accès pour lequel vous souhaitez utiliser le proxy
    const pathname = parse(req.url, true).pathname;
    if (pathname === 'api/user') {
      // Redirigez les requêtes vers le serveur distant
      proxy({
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: { '^/api': '' },
      })(req, res);
    } else {
      // Traitez les autres requêtes avec Next.js
      handle(req, res);
    }
  }).listen(3000, (err: any) => {
    if (err) throw err;
    console.log('> Ready on http://localhost:3000');
  });
});