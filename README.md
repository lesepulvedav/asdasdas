# Evaluación: Refactorización a Infraestructura Modular en AWS

Este repositorio corresponde al **Módulo Raíz (Principal)** que orquesta y conecta de forma dinámica el despliegue de infraestructura en AWS utilizando Terraform. La solución ha sido completamente desacoplada para pasar de un modelo monolítico a un diseño modular de alta escalabilidad.

## 1. Estructura de Módulos y Repositorios

La arquitectura está compuesta por tres componentes independientes distribuidos en repositorios remotos propios, los cuales son llamados desde el archivo `main.tf` central:

* **Módulo de Red (`tarea_terra_networking`):** Provee el entorno de red aislado en la nube.
* **Módulo de Cómputo (`tarea_terra_compute`):** Despliega la lógica del servidor web y su seguridad perimetral.
* **Módulo de Almacenamiento (`tarea_terra_storage`):** Gestiona el almacenamiento estático seguro.

## 2. Descripción de Componentes (Qué hace cada uno)

1. **Networking:** Genera la base de la red mediante una VPC. Define una subred pública con direccionamiento dinámico para el tráfico externo y una subred privada para aislamiento de recursos. Además, configura el Internet Gateway y las tablas de enrutamiento con sus respectivas asociaciones para dar salida a internet.
2. **Compute:** Implementa las reglas de seguridad perimetral a través de un Security Group que permite tráfico a los puertos 22 (SSH) y 80 (HTTP). Dentro de la subred pública, aprovisiona una instancia EC2 configurada mediante `user_data` para automatizar la instalación y el arranque de un servidor web Apache. Consume dinámicamente los IDs de red (VPC y subredes) generados por el módulo anterior.
3. **Storage:** Crea un bucket S3 protegido. Utiliza el recurso `random_id` para garantizar un sufijo único global e impedir colisiones de nombres. Habilita de manera obligatoria el versionamiento de archivos y restringe por completo el acceso público a través de políticas explícitas.

## 3. Justificación del Número de Versión Elegido (SemVer)

Esta entrega ha sido etiquetada formalmente con la versión **`v1.0.0`** debido a los siguientes lineamientos de Semantic Versioning (SemVer):

* **Punto de partida (`v0.1.0`):** Correspondía a la entrega inicial (Prueba 1), la cual consistía en un único archivo monolítico donde todos los recursos estaban acoplados de forma rígida.
* **Consideraciones de las versiones intermedias:** * Crear un módulo nuevo o agregar variables/outputs aumenta el componente **MINOR** (`v0.2.0`).
  * Arreglar un bug o una línea de código específica aumenta el componente **PATCH** (`v0.2.1`).
* **El salto a la versión final (`v1.0.0`):** La separación completa de la lógica en repositorios independientes y la transformación total del código para depender de llamadas a módulos remotos constituye un **cambio de arquitectura estructural**. Dado que el archivo monolítico original se elimina y la nueva forma de despliegue no es compatible hacia atrás con la estructura anterior, las reglas de SemVer exigen aumentar de forma obligatoria el componente **MAJOR**, justificando así la versión final `v1.0.0`.