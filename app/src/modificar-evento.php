<?php
require '../vendor/autoload.php';

use \App\controller\Evento;

if($_SERVER['REQUEST_METHOD'] !== 'POST'){
    header('Content-Type: application/json');
    http_response_code(405);
    echo json_encode([
        "error" => "Método não permitido"
    ]);
    exit();
}

$input = file_get_contents('php://input');

if(empty($input)){
    header('Content-Type: application/json');
    http_response_code(400);
    echo json_encode([
        "error" => "Os dados do evento são obrigatórios"
    ]);
    exit();
}

$data = json_decode($input, true);

$id = $data['id'] ?? null;
$titulo = $data['titulo'] ?? null;
$descricao = $data['descricao'] ?? null;
$dataInicio = $data['inicio'] ?? null;
$dataFim = $data['fim'] ?? null;

if(empty($id) || empty($titulo) || empty($descricao) || empty($dataInicio) || empty($dataFim)){
    header('Content-Type: application/json');
    http_response_code(400);
    echo json_encode([
        "error" => "Todos os campos são obrigatórios"
    ]);
    exit();
}

if(strtotime($dataInicio) > strtotime($dataFim)){
    header('Content-Type: application/json');
    http_response_code(400);
    echo json_encode([
        "error" => "A data de início não pode ser maior que a data de fim"
    ]);
    exit();
}

if(strtotime($dataInicio) < strtotime(date('Y-m-d'))){
    header('Content-Type: application/json');
    http_response_code(400);
    echo json_encode([
        "error" => "A data de início não pode ser menor que a data atual"
    ]);
    exit();
}

$evento = new Evento();
$retorno = $evento->ModificarEvento($id, $titulo, $descricao, $dataInicio, $dataFim);

header('Content-Type: application/json');
echo json_encode([
    "id" => $retorno,
    "mensagem" => 'Evento alterado com sucesso'
]);