resource "docker_image" "grpc_wallarm" {
  name = "awallarm/grpc-wlrm:latest"
}

resource "docker_image" "wallarm_node" {
  name = "wallarm/node:latest"
}

resource "docker_container" "grpc_wallarm" {
  name  = "grpc"
  image = docker_image.grpc_wallarm.latest
  rm = true
}

resource "docker_container" "wallarm_node" {
  name  = "waf"
  image = docker_image.wallarm_node.latest
  rm = true

  env = ["DEPLOY_USER=${var.deploy_user}",
         "DEPLOY_PASSWORD=${var.deploy_password}"]

  ports {
      internal = 80
      external = 5082
  }

  upload {
      file = "/etc/nginx/sites-enabled/default"
      content_base64 = base64encode(<<EOF
      server {
        listen 80 default_server http2;
        server_name localhost;
        root /usr/share/nginx/html;
        index index.html index.htm;
        wallarm_mode monitoring;
	    
        location /ptrav.PathTraversal {
    		grpc_pass grpc://${docker_container.grpc_wallarm.ip_address}:50051;
	    }
}
EOF
)
  }
}