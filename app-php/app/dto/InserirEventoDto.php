<?php
namespace App\dto;

class InserirEventoDto
{
    public $titulo;
    public $descricao;
    public $dataInicio;
    public $dataFim;

    public function __construct($titulo, $descricao, $dataInicio, $dataFim)
    {
        $this->titulo = $titulo;
        $this->descricao = $descricao;
        $this->dataInicio = $dataInicio;
        $this->dataFim = $dataFim;
    }

    public function InserirEvento(){
        $titulo = $this->titulo;
        $descricao = $this->descricao;
        $dataInicio = $this->dataInicio;
        $dataFim = $this->dataFim;
    }

}