package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVPCInfrastructure(t *testing.T) {
	t.Parallel()

	// Obtén el directorio donde se encuentra el código de Terraform
	exampleDir := "../modules/networking"

	// Inicializa el módulo de Terraform
	terraformOptions := &terraform.Options{
		// Dirigirse a la carpeta que contiene los archivos de terraform
		TerraformDir: exampleDir,

		// No generar el archivo de estado
		NoColor: true,

		// Configura variables de entrada si es necesario (puedes añadir más según sea necesario)
		Vars: map[string]interface{}{
			"vpc_name":            "vpc_terra",
			"vpc_cidr":            "10.0.0.0/16",
			"subnets_cidr":        []string{"10.0.1.0/24", "10.0.2.0/24"},
			"availability_zones":  []string{"us-east-1a", "us-east-1b"},
		},

		// Configurar variables de entorno si es necesario (por ejemplo, AWS Access Key)
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": "us-east-1",
		},

		// Configurar variables de backend si es necesario
	}

	// Limpia el estado de Terraform después de que termine el test
	defer terraform.Destroy(t, terraformOptions)

	// Inicializa y aplica el código de Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Verifica que la VPC haya sido creada correctamente
	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	assert.NotEmpty(t, vpcID)

	// Verifica que el Internet Gateway haya sido creado correctamente
	igwID := terraform.Output(t, terraformOptions, "internet_gateway_id")
	assert.NotEmpty(t, igwID)

	// Verifica que las subnets pública y privada estén creadas
	publicSubnetID := terraform.Output(t, terraformOptions, "public_subnet_id")
	assert.NotEmpty(t, publicSubnetID)

	privateSubnetID := terraform.Output(t, terraformOptions, "private_subnet_id")
	assert.NotEmpty(t, privateSubnetID)

	// Verifica que la tabla de rutas pública exista
	routeTableID := terraform.Output(t, terraformOptions, "route_table_id")
	assert.NotEmpty(t, routeTableID)
}