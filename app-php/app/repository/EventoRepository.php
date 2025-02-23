<?php
namespace App\repository;

use App\dto\InserirEventoDto;
use App\dto\DeletarEventoDto;
use App\dto\ModificarEventoDto;
use App\dto\RetornaEventoDto;
use PDO;
use PDOException;
use Dotenv\Dotenv;

class EventoRepository
{
    public function __construct()
    {
        $dotenv = Dotenv::createImmutable(__DIR__ . '/../../');
        $dotenv->load();
    }
    public function InserirEvento(InserirEventoDto $dto)
    {
        $titulo = $dto->titulo;
        $descricao = $dto->descricao;
        $dataInicio = $dto->dataInicio;
        $dataFim = $dto->dataFim;

        $dataInicio = $this->converteDataParaMysql($dataInicio);
        $dataFim = $this->converteDataParaMysql($dataFim);

        $sql = "INSERT INTO eventos (titulo, descricao, inicio, fim) VALUES ('$titulo', '$descricao', '$dataInicio', '$dataFim')";

        try {
            $conexao = $this->retornaConexao();
            $conexao->exec($sql);
            return $conexao->lastInsertId();
        } catch (PDOException $e) {
            header("Content-Type: application/json");
            http_response_code(500);
            echo json_encode([
                "error" =>  "Erro ao executar a consulta ao banco de dados: " . $e->getMessage(),
            ]);
            exit();
        }
    }

    public function ListarEventos()
    {
        $sql = "SELECT * FROM eventos WHERE deleted_at IS NULL ORDER BY inicio";

        try {
            $conexao = $this->retornaConexao();
            $resultado = $conexao->query($sql);

            return $resultado->fetchAll(\PDO::FETCH_ASSOC);
        } catch (PDOException $e) {
            header("Content-Type: application/json");
            http_response_code(500);
            echo json_encode([
                "error" =>
                    "Erro ao executar a consulta ao banco de dados: " . $e->getMessage(),
            ]);
            exit();
        }
        
    }

    public function RetornaEvento(RetornaEventoDto $dto)
    {
        $id = $dto->id;

        if (!$this->existeEvento($id)) {
            header("Content-Type: application/json");
            http_response_code(404);
            echo json_encode([
                "error" => "Evento não encontrado",
            ]);
            exit();
        }

        $sql = "SELECT * FROM eventos WHERE id = ? AND deleted_at IS NULL";

        try {
            $conexao = $this->retornaConexao();
            $stmt = $conexao->prepare($sql);
            $stmt->execute([$id]);
            return $stmt->fetch(\PDO::FETCH_ASSOC);
        } catch (PDOException $e) {
            header("Content-Type: application/json");
            http_response_code(500);
            echo json_encode([
                "error" => "Erro ao executar a consulta ao banco de dados: " . $e->getMessage(),
            ]);
            exit();
        }
    }

    public function DeletarEvento(DeletarEventoDto $dto)
    {
        $id = $dto->id;

        if (!$this->existeEvento($id)) {
            header("Content-Type: application/json");
            http_response_code(404);
            echo json_encode([
                "error" => "Evento não encontrado",
            ]);
            exit();
        }

        $sql = "UPDATE eventos SET deleted_at = NOW() WHERE id = ?";

        try {
            $conexao = $this->retornaConexao();
            $stmt = $conexao->prepare($sql);
            $stmt->execute([$id]);
            return $stmt->rowCount();
        } catch (PDOException $e) {
            header("Content-Type: application/json");
            http_response_code(500);
            echo json_encode([
                "error" => "Erro ao executar a consulta ao banco de dados: " . $e->getMessage(),
            ]);
            exit();
        }
    }

    public function ModificarEvento(ModificarEventoDto $dto)
    {
        $id = $dto->id;
        $titulo = $dto->titulo;
        $descricao = $dto->descricao;
        $dataInicio = $dto->dataInicio;
        $dataFim = $dto->dataFim;

        if (!$this->existeEvento($id)) {
            header("Content-Type: application/json");
            http_response_code(404);
            echo json_encode([
                "error" => "Evento não encontrado",
            ]);
            exit();
        }

        $dataInicio = $this->converteDataParaMysql($dataInicio);
        $dataFim = $this->converteDataParaMysql($dataFim);

        $sql = "UPDATE eventos SET titulo = ?, descricao = ?, inicio = ?, fim = ? WHERE id = ?";

        try {
            $conexao = $this->retornaConexao();
            $stmt = $conexao->prepare($sql);
            $stmt->execute([$titulo, $descricao, $dataInicio, $dataFim, $id]);
            return $stmt->rowCount();
        } catch (PDOException $e) {
            header("Content-Type: application/json");
            http_response_code(500);
            echo json_encode([
                "error" => "Erro ao executar a consulta ao banco de dados: " . $e->getMessage(),
            ]);
            exit();
        }
    }

    private function retornaConexao()
    {
        
        $DB_HOST = $_ENV['DB_HOST'];
        $DB_USER = $_ENV['DB_USER'];
        $DB_PASS = $_ENV['DB_PASS'];
        $DB_NAME = $_ENV['DB_NAME'];

        try {
            $conn = new \PDO("mysql:host=$DB_HOST;dbname=$DB_NAME", $DB_USER, $DB_PASS);
                $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
                return $conn;
            } catch(PDOException $e) {
                header("Content-Type: application/json");
                http_response_code(500);
                echo json_encode([
                    "error" =>
                        "Erro ao conectar ao banco de dados: " . $e->getMessage(),
                ]);
                exit();
            }
    }

    private function converteDataParaMysql($data)
    {
        return date('Y-m-d h:i:s', strtotime($data));
    }

    private function existeEvento($id)
    {
        $sql = "SELECT * FROM eventos WHERE id = ? AND deleted_at IS NULL";

        try{
            $conexao = $this->retornaConexao();
            $stmt = $conexao->prepare($sql);
            $stmt->execute([$id]);
            return $stmt->rowCount() > 0;
        } catch (PDOException $e) {
            header("Content-Type: application/json");
            http_response_code(500);
            echo json_encode([
                "error" => "Erro ao executar a consulta ao banco de dados: " . $e->getMessage(),
            ]);
            exit();
        }
    }
}