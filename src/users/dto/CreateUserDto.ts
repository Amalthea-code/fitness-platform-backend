type BaseCreateUserDto = {
    email: string
    password: string
}

type CreateTrainerDto = BaseCreateUserDto & {
    role: 'TRAINER'
    specialization?: string
    experience?: number
}

type CreateClientDto = BaseCreateUserDto & {
    role: 'CLIENT'
    weight?: number
    height?: number
}

type CreateAdminDto = BaseCreateUserDto & {
    role: 'ADMIN'
}

export type CreateUserDto =
    | CreateTrainerDto
    | CreateClientDto
    | CreateAdminDto
