package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func GenerateExcelReport(c *gin.Context) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BASE_URL não definida"})
		return
	}

	reqRank, _ := http.NewRequest("GET", baseURL+"/services/rank-clients", nil)
	reqRank.Header.Set("Authorization", c.GetHeader("Authorization"))
	rankResp, err := http.DefaultClient.Do(reqRank)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar rank-clients"})
		return
	}
	defer rankResp.Body.Close()

	reqOrder, _ := http.NewRequest("GET", baseURL+"/services/ordes-in-progress", nil)
	reqOrder.Header.Set("Authorization", c.GetHeader("Authorization"))
	orderResp, err := http.DefaultClient.Do(reqOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar orders-in-progress"})
		return
	}
	defer orderResp.Body.Close()

	reqSummary, _ := http.NewRequest("GET", baseURL+"/services/summary", nil)
	reqSummary.Header.Set("Authorization", c.GetHeader("Authorization"))
	summaryResp, err := http.DefaultClient.Do(reqSummary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar summary"})
		return
	}
	defer summaryResp.Body.Close()

	var rankings []map[string]interface{}
	bodyRank, _ := ioutil.ReadAll(rankResp.Body)
	json.Unmarshal(bodyRank, &rankings)

	var orders []map[string]interface{}
	bodyOrder, _ := ioutil.ReadAll(orderResp.Body)
	json.Unmarshal(bodyOrder, &orders)

	var summary map[string]interface{}
	bodySummary, _ := ioutil.ReadAll(summaryResp.Body)
	json.Unmarshal(bodySummary, &summary)

	f := excelize.NewFile()

	index1, err := f.NewSheet("RankClients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar aba RankClients"})
		return
	}
	f.SetCellValue("RankClients", "A1", "Email")
	f.SetCellValue("RankClients", "B1", "Criado em")
	f.SetCellValue("RankClients", "C1", "Quantidade de Pedidos")

	for i, client := range rankings {
		f.SetCellValue("RankClients", fmt.Sprintf("A%d", i+2), client["email"])
		f.SetCellValue("RankClients", fmt.Sprintf("B%d", i+2), client["createdAt"])
		f.SetCellValue("RankClients", fmt.Sprintf("C%d", i+2), client["ordersQuantity"])
	}

	f.NewSheet("OrdersInProgress")
	f.SetCellValue("OrdersInProgress", "A1", "ID Pedido")
	f.SetCellValue("OrdersInProgress", "B1", "Email Cliente")
	f.SetCellValue("OrdersInProgress", "C1", "Produtos")
	f.SetCellValue("OrdersInProgress", "D1", "Status")
	f.SetCellValue("OrdersInProgress", "E1", "Preço Total")

	for i, order := range orders {
		productsStr := ""
		if products, ok := order["products"].([]interface{}); ok {
			for _, p := range products {
				if productMap, ok := p.(map[string]interface{}); ok {
					name := productMap["name"]
					qtd := productMap["quantity"]
					productsStr += fmt.Sprintf("%v (Qtd: %v), ", name, qtd)
				}
			}
		}

		client := order["client"].(map[string]interface{})
		f.SetCellValue("OrdersInProgress", fmt.Sprintf("A%d", i+2), order["id"])
		f.SetCellValue("OrdersInProgress", fmt.Sprintf("B%d", i+2), client["email"])
		f.SetCellValue("OrdersInProgress", fmt.Sprintf("C%d", i+2), productsStr)
		f.SetCellValue("OrdersInProgress", fmt.Sprintf("D%d", i+2), order["status"])
		f.SetCellValue("OrdersInProgress", fmt.Sprintf("E%d", i+2), order["totalPrice"])
	}

	f.NewSheet("Summary")
	f.SetCellValue("Summary", "A1", "Total de Vendas")
	f.SetCellValue("Summary", "B1", "Montante Faturado")
	f.SetCellValue("Summary", "C1", "Total de Pedidos")

	f.SetCellValue("Summary", "A2", summary["totalSalesMade"])
	f.SetCellValue("Summary", "B2", summary["invoicedSmount"])
	f.SetCellValue("Summary", "C2", summary["totalOrdersSold"])

	f.SetActiveSheet(index1)
	f.DeleteSheet("Sheet1")

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar relatório"})
		return
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", `attachment; filename="relatorio-gestao-de-vendas.xlsx"`)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}
