package injector

import (
	repoAdmin "kang-sayur-backend/infrastructure/repository/admin"
	repoAddress "kang-sayur-backend/infrastructure/repository/adress"
	repoCart "kang-sayur-backend/infrastructure/repository/cart"
	repoCategory "kang-sayur-backend/infrastructure/repository/category"
	repoCustomer "kang-sayur-backend/infrastructure/repository/customer"
	repoFeedback "kang-sayur-backend/infrastructure/repository/feedback"
	repoForgotPassword "kang-sayur-backend/infrastructure/repository/forgot_password"
	repoGrocery "kang-sayur-backend/infrastructure/repository/grocery"
	repoGroceryImage "kang-sayur-backend/infrastructure/repository/grocery_image"
	repoInvoice "kang-sayur-backend/infrastructure/repository/invoice"
	repoInvoiceItem "kang-sayur-backend/infrastructure/repository/invoice_item"
	repoPermission "kang-sayur-backend/infrastructure/repository/permission"
	repoRecipe "kang-sayur-backend/infrastructure/repository/recipe"
	repoRecipeDetail "kang-sayur-backend/infrastructure/repository/recipe_detail"
	repoRecipeImage "kang-sayur-backend/infrastructure/repository/recipe_image"
	repoSubAdmin "kang-sayur-backend/infrastructure/repository/sub_admin"
	repoVerification "kang-sayur-backend/infrastructure/repository/verification"

	"kang-sayur-backend/model/domain/address"
	"kang-sayur-backend/model/domain/admin"
	"kang-sayur-backend/model/domain/cart"
	"kang-sayur-backend/model/domain/category"
	customer "kang-sayur-backend/model/domain/customer"
	"kang-sayur-backend/model/domain/feedback"
	forgotpassword "kang-sayur-backend/model/domain/forgot_password"
	"kang-sayur-backend/model/domain/grocery"
	groceryimage "kang-sayur-backend/model/domain/grocery_image"
	invoice "kang-sayur-backend/model/domain/invoice"
	invoiceitem "kang-sayur-backend/model/domain/invoice_item"
	permission "kang-sayur-backend/model/domain/permission"
	recipe "kang-sayur-backend/model/domain/recipe"
	recipedetail "kang-sayur-backend/model/domain/recipe_detail"
	recipeimage "kang-sayur-backend/model/domain/recipe_images"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	verification "kang-sayur-backend/model/domain/verification"
)

type repositoriesList struct {
	Address        address.AddressRepository
	Admin          admin.AdminRepository
	Cart           cart.CartRepository
	Category       category.CategoryRepository
	Customer       customer.CustomerRepository
	Feedback       feedback.FeedbackRepository
	ForgotPassword forgotpassword.ForgotPasswordRepository
	Grocery        grocery.GroceryRepository
	GroceryImage   groceryimage.GroceryImageRepository
	Invoice        invoice.InvoiceRepository
	InvoiceItem    invoiceitem.InvoiceItemRepository
	Permission     permission.PermissionRepository
	Recipe         recipe.RecipeRepository
	RecipeDetail   recipedetail.RecipeDetailRepository
	RecipeImage    recipeimage.RecipeImagesRepository
	SubAdmin       subadmin.SubAdminRepository
	Verification   verification.VerificationRepository
}

// private access only!
func (sp *Infrastructures) Repositories() *repositoriesList {
	db := sp.DatabaseMongodb()

	return &repositoriesList{
		Address:        repoAddress.AddressRepository(db),
		Admin:          repoAdmin.AdminRepository(db),
		Cart:           repoCart.CartRepository(db),
		Category:       repoCategory.CategoryRepository(db),
		Customer:       repoCustomer.CustomerRepository(db),
		Feedback:       repoFeedback.FeedbackRepository(db),
		ForgotPassword: repoForgotPassword.ForgotPassword(db),
		Grocery:        repoGrocery.GroceryRepository(db),
		GroceryImage:   repoGroceryImage.GroceryImageRepository(db),
		Invoice:        repoInvoice.InvoiceRepository(db),
		InvoiceItem:    repoInvoiceItem.InvoiceItemRepository(db),
		Permission:     repoPermission.PermissionRepository(db),
		Recipe:         repoRecipe.RecipeRepository(db),
		RecipeDetail:   repoRecipeDetail.RecipeDetailRepository(db),
		RecipeImage:    repoRecipeImage.RecipeImageRepository(db),
		SubAdmin:       repoSubAdmin.SubAdminRepository(db),
		Verification:   repoVerification.VerificationRepository(db),
	}
}
