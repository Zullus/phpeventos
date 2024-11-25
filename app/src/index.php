<?php
$rotasExistentes = [
    '/src/adiciona-evento' => '/adiciona-evento.php',
    '/src/lista-eventos' => '/lista-eventos.php',
    '/src/deletar-evento' => '/deletar-evento.php',
    '/src/modificar-evento' => '/modificar-evento.php',
    '/src/retornar-evento' => '/retornar-evento.php',
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
