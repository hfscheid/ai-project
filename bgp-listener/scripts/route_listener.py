from sys import stdout
import socketserver
import http.server 
import cgi 

PORT = 5555

class RequestHandler(http.server.SimpleHTTPRequestHandler):

    def createResponse(self, command):
        self.send_response(200)
        self.send_header('Content-Type', 'application/text')
        self.end_headers()
        self.wfile.write(bytes(command, 'utf-8'))

    def do_POST(self):

        #receive message and then stdout
        form = cgi.FieldStorage(
            fp=self.rfile,
            headers=self.headers,
            environ={'REQUEST_METHOD':'POST'})
        command = form.getvalue('command')
        stdout.write('%s\n' % command)
        stdout.flush()
        self.createResponse('Success: %s\n' % command)


handler = RequestHandler
httpd = socketserver.TCPServer(('', PORT), handler)
stdout.flush()
httpd.serve_forever()