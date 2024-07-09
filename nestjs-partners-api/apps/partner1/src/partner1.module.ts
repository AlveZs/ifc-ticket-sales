import { Module } from '@nestjs/common';
import { SpotsModule } from './spots/spots.module';
import { EventsModule } from './events/events.module';

@Module({
  imports: [EventsModule, SpotsModule],
})
export class Partner1Module {}
