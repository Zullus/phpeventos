<?php
require '../vendor/autoload.php';

use \App\controller\Evento;

$id = $_GET['id'];

if (empty($id)) {
    header("Content-Type: application/json");
    http_response_code(400);
    echo json_encode([
        "error" => "O id do evento é obrigatório",
    ]);
    exit();
}

if (empty($id) || !is_numeric($id)) {
    header("Content-Type: application/json");
    http_response_code(400);
    echo json_encode([
        "error" => "O id do evento é obrigatório e deve ser um número",
    ]);
    exit();
}

$evento = new Evento();
$evento->DeletarEvento($id);

header("Content-Type: application/json");
http_response_code(200);
echo json_encode([
    "message" => "Evento apagado com sucesso",
    "success" => true,
]);