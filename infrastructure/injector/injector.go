package injector

import (
	mailer "kang-sayur-backend/infrastructure/SMTP"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/identifier"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/model/domain/address"
	"kang-sayur-backend/model/domain/admin"
	"kang-sayur-backend/model/domain/cart"
	"kang-sayur-backend/model/domain/category"
	customer "kang-sayur-backend/model/domain/customer"
	"kang-sayur-backend/model/domain/domicile"
	"kang-sayur-backend/model/domain/feedback"
	forgotpassword "kang-sayur-backend/model/domain/forgot_password"
	"kang-sayur-backend/model/domain/grocery"
	groceryimage "kang-sayur-backend/model/domain/grocery_image"
	groceryprice "kang-sayur-backend/model/domain/grocery_price"
	invoice "kang-sayur-backend/model/domain/invoice"
	invoiceitem "kang-sayur-backend/model/domain/invoice_item"
	permission "kang-sayur-backend/model/domain/permission"
	recipe "kang-sayur-backend/model/domain/recipe"
	recipedetail "kang-sayur-backend/model/domain/recipe_detail"
	recipeimage "kang-sayur-backend/model/domain/recipe_images"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	verification "kang-sayur-backend/model/domain/verification"
)

type injectorProvider struct {
	Encryption encryption.EncryptionContract
	Identifier identifier.IdentifierContract
	JWT        jsonwebtoken.JWTContract
	SMTP       mailer.Contract
}

type repositoryProvider struct {
	Address        address.AddressRepository
	Admin          admin.AdminRepository
	Cart           cart.CartRepository
	Category       category.CategoryRepository
	Customer       customer.CustomerRepository
	Domicile       domicile.DomicileRepository
	Feedback       feedback.FeedbackRepository
	ForgotPassword forgotpassword.ForgotPasswordRepository
	Grocery        grocery.GroceryRepository
	GroceryImage   groceryimage.GroceryImageRepository
	GroceryPrice   groceryprice.PriceRepository
	Invoice        invoice.InvoiceRepository
	InvoiceItem    invoiceitem.InvoiceRepository
	Permission     permission.PermissionRepository
	Recipe         recipe.RecipeRepository
	RecipeDetail   recipedetail.RecipeDetailRepository
	RecipeImage    recipeimage.RecipeImagesRepository
	SubAdmin       subadmin.SubAdminRepository
	Verification   verification.VerificationRepository
}
