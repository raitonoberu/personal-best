package handler

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

const maxDocumentSize = 1024 * 1024 * 10 // 10 MB

var allowedContentTypes = []string{
	"image/jpeg", "image/png",
	"image/webp", "image/gif",
	"application/pdf",
}

// @Summary Save document
// @Security Bearer
// @Description Save player document.
// @Description File can be one of:
// @Description jpeg, png, webp, gif, pdf
// @Tags document
// @Accept mpfd
// @Produce json
// @Param document formData file true "document"
// @Param name formData string true "name"
// @Success 204
// @Router /api/documents [post]
func (h Handler) SaveDocument(c echo.Context) error {
	document, err := c.FormFile("document")
	if err != nil {
		return err
	}
	if document.Size > maxDocumentSize {
		return ErrFileTooBig
	}

	contentType := document.Header.Get("content-type")
	if contentType == "" {
		contentType = guessContentType(document.Filename)
	}
	if !isAllowed(contentType) {
		return ErrFileUnsupported
	}

	name := c.FormValue("name")
	if name == "" {
		name = document.Filename
	}

	file, err := document.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	err = h.service.SaveDocument(c.Request().Context(),
		getUserID(c), file, name, document.Size, contentType)
	if err != nil {
		return err
	}
	return c.NoContent(204)
}

// @Summary List player documents
// @Security Bearer
// @Description List documents of player
// @Tags document
// @Produce json
// @Param user_id path int true "id of user"
// @Success 200 {object} []model.Document
// @Router /api/users/{user_id}/documents [get]
func (h Handler) ListDocuments(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if id != getUserID(c) && !h.getUserRole(c).IsAdmin {
		return ErrAccessDenied
	}

	documents, err := h.service.ListDocuments(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(200, documents)
}

// @Summary Get document
// @Security Bearer
// @Description Get URL of document
// @Tags document
// @Produce json
// @Param id path int true "id of document"
// @Success 200 {object} model.GetDocumentResponse
// @Router /api/documents/{id} [get]
func (h Handler) GetDocument(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	url, err := h.service.GetDocument(c.Request().Context(),
		id, getUserID(c), h.getUserRole(c).IsAdmin)
	if err != nil {
		return err
	}
	return c.JSON(200, model.GetDocumentResponse{
		URL: url,
	})
}

func guessContentType(filename string) string {
	parts := strings.Split(filename, ".")
	ext := parts[len(parts)-1]

	switch strings.ToLower(ext) {
	case "jpg", "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "webp":
		return "image/webp"
	case "gif":
		return "image/gif"
	case "pdf":
		return "application/pdf"
	default:
		return "unknown"
	}
}

func isAllowed(contentType string) bool {
	for _, t := range allowedContentTypes {
		if t == contentType {
			return true
		}
	}
	return false
}
