import { Injectable } from '@nestjs/common';

@Injectable()
export class PlansService {
    constructor(private prisma: PlansService) {}
}
