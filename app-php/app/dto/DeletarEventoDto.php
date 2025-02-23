<?php
namespace App\dto;

class DeletarEventoDto
{
    public $id;

    public function __construct($id)
    {
        $this->id = $id;
    }

    public function DeletarEvento(){
        $id = $this->id;
    }

}