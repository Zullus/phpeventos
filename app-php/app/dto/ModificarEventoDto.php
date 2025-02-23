<?php

namespace App\dto;

class ModificarEventoDto
{
    public $id;
    public $titulo;
    public $descricao;
    public $dataInicio;
    public $dataFim;

    public function __construct($id, $titulo, $descricao, $dataInicio, $dataFim)
    {
        $this->id = $id;
        $this->titulo = $titulo;
        $this->descricao = $descricao;
        $this->dataInicio = $dataInicio;
        $this->dataFim = $dataFim;
    }

    public function ModificarEvento()
    {
        $id = $this->id;
        $titulo = $this->titulo;
        $descricao = $this->descricao;
        $dataInicio = $this->dataInicio;
        $dataFim = $this->dataFim;
    }
}
