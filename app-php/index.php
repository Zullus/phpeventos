<?php
ini_set('display_errors',1);
ini_set('display_startup_erros',1);
error_reporting(E_ALL);

require './vendor/autoload.php';

$rotasExistentes = [
    '/adiciona-evento' => '/app/adiciona-evento.php',
    '/lista-eventos' => '/app/lista-eventos.php',
    '/deletar-evento' => '/app/deletar-evento.php',
    '/modificar-evento' => '/app/modificar-evento.php',
    '/retorna-evento' => '/app/retorna-evento.php',
];

$rota = $_SERVER['REQUEST_URI'];
$rota = explode('?', $rota)[0];

if (array_key_exists($rota, $rotasExistentes)) {
    require __DIR__ . $rotasExistentes[$rota];
} else {
    header('Content-Type: application/json');
    http_response_code(404);
    echo json_encode(['erro' => 'Rota não encontrada']);
}
