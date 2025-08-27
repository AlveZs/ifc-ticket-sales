import { PropsWithChildren } from "react";
import { EventModel } from "../models";
import Link from "next/link";
import { EventImage } from "./EventImage";

export type EventCardProps = {
  event: EventModel;
};

export function EventCard(props: PropsWithChildren<EventCardProps>) {
  return (
    <Link href={`/event/${props.event.id}/spots-layout`}>
      <div className="flex w-[277px] h-full flex-col rounded-2xl bg-secondary">
        <EventImage src={props.event.image_url} alt={props.event.name} />
        <div className="flex flex-col gap-y-2 px-4 py-6" >
          <p className="text-sm uppercase text-subtitle">
            {new Date(props.event.date).toLocaleDateString("en-US", {
              weekday: "long",
              day: "2-digit",
              year: "numeric",
            })}
          </p>
          <p className="font-semibold">{props.event.name}</p>
          <p className="text-sm font-normal" >{props.event.location}</p>
        </div>
      </div>
    </Link>
  );
}
