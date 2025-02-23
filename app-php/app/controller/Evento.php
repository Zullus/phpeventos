<?php

namespace App\controller;

use App\repository\EventoRepository;
use App\dto\InserirEventoDto;
use App\dto\DeletarEventoDto;
use App\dto\ModificarEventoDto;
use App\dto\RetornaEventoDto;

class Evento
{
    public function InserirEvento($titulo, $descricao, $dataInicio, $dataFim)
    {
        $evento = new InserirEventoDto($titulo, $descricao, $dataInicio, $dataFim);
        $repository = new EventoRepository();
        return $repository->InserirEvento($evento);
    }

    public function ListarEventos()
    {
        $repository = new EventoRepository();
        return $repository->ListarEventos();
    }

    public function DeletarEvento($id)
    {
        $evento = new DeletarEventoDto($id);
        $repository = new EventoRepository();
        return $repository->DeletarEvento($evento);
    }

    public function ModificarEvento($id, $titulo, $descricao, $dataInicio, $dataFim)
    {
        $evento = new ModificarEventoDto($id, $titulo, $descricao, $dataInicio, $dataFim);
        $repository = new EventoRepository();
        return $repository->ModificarEvento($evento);
    }

    public function RetornaEvento($id)
    {
        $evento = new RetornaEventoDto($id);
        $repository = new EventoRepository();
        return $repository->RetornaEvento($evento);
    }
}