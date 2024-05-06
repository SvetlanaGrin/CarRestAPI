package handler

import (
	"carRestAPI/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// @Summary Create
// @Tags createCar
// @Description Get data from the machine catalog from an external API
// @Accept  application/json
// @Produce  application/json
// @Param input body models.Car true "data car"
// @Success 200 {object}      models.Answer
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/ [post]
func (h *Handler) Create(c *gin.Context) {
	var regNums models.GetRegNums

	
	if err := c.BindJSON(&regNums); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())  
		return
	}
	logrus.Debug(fmt.Sprintf("request body decoded.RegNums %s", regNums))

	input := make([]models.Car, len(regNums.RegNums))
	for i, elem := range regNums.RegNums {
		url := fmt.Sprintf("https://89c38835-9610-458b-8ee6-a0e3fc0bc709.mock.pstmn.io//info?regNum=%s", elem)

		resp, err := http.Get(url)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&input[i]); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	err := h.services.RequestCarCatalog.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var ans models.Answer
	ans.Status="Ok"
	c.JSON(http.StatusOK, ans)
}
// @Summary Delete
// @Tags DeleteCar
// @Description  Delete data from the machine catalog
// @Accept  application/json
// @Produce  application/json
// @Param regNum  path string true "regNum"
// @Success 200 {object}    models.Answer
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router	/api/{regNum} [delete]
func (h *Handler) Delete(c *gin.Context) {
	regNum := c.Param("regNum")
	logrus.Debug(fmt.Sprintf("request paranm decoded.RegNum %s", regNum))

	err := h.services.RequestCarCatalog.Delete(regNum)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var ans models.Answer
	ans.Status="Ok"
	c.JSON(http.StatusOK, ans)
}
// @Summary Update
// @Tags UpdateCar
// @Description  Update data from the machine catalog
// @Accept  application/json
// @Produce  application/json
// @Param regNum  path string true "regNum"
// @Param input body models.Car true "data car"
// @Success 200 {object}    models.Answer
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router	/api/{regNum} [put]
func (h *Handler) Update(c *gin.Context) {
	regNum := c.Param("regNum")
	logrus.Debug(fmt.Sprintf("request paranm decoded.RegNum %s", regNum))

	var input models.UpdateCar
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.RequestCarCatalog.Update(regNum, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	var ans models.Answer
	ans.Status="Ok"
	c.JSON(http.StatusOK, ans)

}


// @Summary GetAll
// @Tags GetAllCar
// @Description  Get all data from the machine catalog
// @Accept  application/json
// @Produce  application/json
// @Param mark query string false "mark"
// @Param limit  query int false  "limit"
// @Param offset  query int false  "offset"
// @Success 200 {object}    []models.Car
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router	/api/ [get]
func (h *Handler) GetAll(c *gin.Context) {

	mark := c.Query("mark")
	logrus.Debug(fmt.Sprintf("request query fltration decoded.mark %s", mark))
	params := NewParams(c)
	logrus.Debug(fmt.Sprintf("request query pfgination decoded. limit=%s and offset=%s", strconv.Itoa(params.Limit), strconv.Itoa(params.Offset)))
	cars, err := h.services.RequestCarCatalog.GetAll(params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.GetAllCarResponse{
		Data: cars,
	})

}
