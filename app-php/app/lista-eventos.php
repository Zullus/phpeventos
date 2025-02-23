<?php

use \App\controller\Evento;

$evento = new Evento();
$eventos = $evento->ListarEventos();

header('Content-Type: application/json');
http_response_code(200);
echo json_encode($eventos);
