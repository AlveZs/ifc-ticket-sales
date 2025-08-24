import { SpotSeat } from "@/app/components/SpotSeat";
import { Title } from "@/app/components/Title";
import { EventModel, SpotModel } from "@/app/models"
import Link from "next/link";

export default function SpotsLayoutPage() {
  const event: EventModel = {
    id: "1",
    name: "Test",
    organization: "Productions",
    date: "2022-12-31T00:00:00.000Z",
    location: "Salvador",
  };
  const spot: SpotModel = {
    name: "A1",
  }

  return (
    <main className="mt-10">
      <div className="flex w-[1176px] max-w-full flex-row flex-wrap justify-center gap-x-8 rounded-2xl bg-secondary p-4 md:justify-normal">
        <img src="/image.png" alt="" />
        <div className="flex max-w-full flex-col gap-y-6">
          <div className="flex flex-col gap-y-2">
            <p className="text-sm font-semibold uppercase text-subtitle">
              {new Date(event.date).toLocaleDateString("pt-BR", {
                weekday: "long",
                day: "2-digit",
                month: "2-digit",
                year: "numeric",
              })}
            </p>
            <p className="text-2xl font-semibold">{event.name}</p>
            <p className="font-normal">{event.location}</p>
          </div>
          <div className="flex h-[128px] flex-wrap justify-between gap-y-5 gap-x-3">
            <div className="flex flex-col gap-y-2">
              <p className="font-semibold">Organizer</p>
              <p className="text-sm font-normal">{event.organization}</p>
            </div>
            <div className="flex flex-col gap-y-2">
              <p className="font-semibold">Rating</p>
              <p className="text-sm font-normal">{event.rating}</p>
            </div>
          </div>
        </div>
      </div>
      <Title className="mt-10">Choose your spot</Title>
      <div className="mt-6 flex flex-wrap justify-between">
        <div className="mb-4 flex w-full max-w-[650px] flex-col gap-y-8 rounded-2xl bg-secondary p-6">
          <div className="rounded-2xl bg-bar py-4 text-center text-[20px] font-bold uppercase text-white">
            Stage
          </div>
          <div className="md:w-full md:justify-normal">
            <div className="flex flex-row gap-3 items-center mb-3">
              <div className="w-4">A</div>
              <div className="ml-2 flex flex-row">
                <SpotSeat
                  key={spot.name}
                  spotId={spot.name}
                  spotLabel={spot.name.slice(1)}
                  reserved={false}
                  disabled={false}
                />
              </div>
            </div>
          </div>
          <div className="flex w-full flex-row justify-around">
              <div className="flex flex-row items-center">
                <span className="mr-1 block h-4 w-4 rounded-full bg-[#00A96E]"></span>
                Available
              </div>
              <div className="flex flex-row items-center">
                <span className="mr-1 block h-4 w-4 rounded-full bg-[#A6ADBB]"></span>
                Occupied
              </div>
              <div className="flex flex-row items-center">
                <span className="mr-1 block h-4 w-4 rounded-full bg-(--blurple)"></span>
                Selected
              </div>
          </div>
        </div>
        <div className="flex w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary px-4 py-6">
          <h1 className="text-[20px] font-semibold">
            Check out the event prices
          </h1>
          <p>
            Full: {"$ 100.00"} <br />
            Half price: {`$ 50.00`}
          </p>
          <div className="flex flex-col">
            Select ticket type option
          </div>
          <div>Total: {`$ total price`}</div>
          <Link
            href="/checkout"
            className="rounded-lg bg-btn-primary py-5 text-sm font-semibold uppercase text-btn-primary text-center hover:bg-white"
          >
              Go to checkout
          </Link>
        </div>
      </div>
    </main>
  );
}