import { EventCard } from "./components/EventCard";
import { Title } from "./components/Title";
import { EventModel } from "./models";

export default function HomePage() {
  const events: EventModel[] = [
    {
      id: "1",
      name: "Test",
      organization: "Productions",
      date: "2022-12-31T00:00:00.000Z",
      location: "Salvador",
    },
    {
      id: "2",
      name: "Test",
      organization: "Productions",
      date: "2022-12-31T00:00:00.000Z",
      location: "Salvador",
    },
    {
      id: "3",
      name: "Test",
      organization: "Productions",
      date: "2022-12-31T00:00:00.000Z",
      location: "Salvador",
    }
  ];
  return (
    <main className="mt-10 flex flex-col">
      <Title>Available Events</Title>
      <div className="mt-8 sm:grid sm:grid-cols-(--auto-fit-cards) flex flex-wrap justify-center gap-x-2 gap-y-4">
        {events.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>
    </main>
  );
}
