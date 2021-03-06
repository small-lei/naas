package dao

var (
	OAuth2Client         OAuth2Clienter        = &oauth2Client{}
	OAuth2ClientInfo     OAuth2ClientInfoer    = &oauth2ClientInfo{}
	OAuth2ClientScope    OAuth2ClientScoper    = &oauth2ClientScope{}
	OAuth2Scope          OAuth2Scoper          = &oauth2Scope{}
	Resource             Resourcer             = &resource{}
	ResourceWebRoute     ResourceWebRouter     = &resourceWebRoute{}
	Admin                Adminer               = &admin{}
	User                 Userer                = &user{}
	UserInfo             UserInfoer            = &userInfo{}
	Organization         Organizationer        = &organization{}
	OrganizationRole     OrganizationRoleer    = &organizationRole{}
	Role                 Roleer                = &role{}
	UserRole             UserRoleer            = &userRole{}
	RoleResourceWebRoute RoleResourceWebRouter = &roleResourceWebRoute{}
)
