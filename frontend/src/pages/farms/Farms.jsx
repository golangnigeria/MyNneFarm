import React from "react";
import { useQuery } from "@tanstack/react-query";
import { fetchFarms } from "../../api/farms";
import { Leaf, MapPin, TrendingUp, PackageCheck } from "lucide-react";

function Farms() {
  const {
    data: farms,
    isLoading,
    isError,
  } = useQuery({
    queryKey: ["farms"],
    queryFn: fetchFarms,
  });

  const defaultImage =
    "https://images.unsplash.com/photo-1580327331400-dc0c0be4a79c?auto=format&fit=crop&w=800&q=60";

  if (isLoading) {
    return (
      <div className="text-center py-10 text-[#1F3B17]">Loading farms...</div>
    );
  }

  if (isError) {
    return (
      <div className="text-center py-10 text-red-700">
        Failed to load farms. Please try again later.
      </div>
    );
  }

  return (
    <section className="bg-[#F3E7D5] min-h-screen px-4 py-16">
      <h1 className="text-3xl font-bold text-center text-[#1F3B17] mb-10">
        Available Farms
      </h1>

      {farms && farms.length > 0 ? (
        <div className="flex flex-col gap-6">
          {farms.map((farm) => {
            const unitsLeft = farm.units_available - (farm.units_sold || 0);
            const totalPotentialRevenue =
              farm.price_per_unit * farm.units_available;
            const expectedReturn = totalPotentialRevenue * farm.expected_roi;

            return (
              <div
                key={farm.id}
                className="bg-white/40 backdrop-blur-md border border-white/30 rounded-2xl shadow-md p-4 flex flex-col md:flex-row gap-4 transition-transform hover:scale-[1.015] max-w-full md:max-w-[90%] md:mx-auto"
              >
                {/* Image */}
                <img
                  src={farm.image_url || defaultImage}
                  alt={farm.title}
                  className="w-full md:w-1/2 h-48 md:h-auto object-cover rounded-xl"
                />

                {/* Info Section */}
                <div className="flex flex-col justify-between flex-1">
                  <div>
                    <h2 className="text-2xl font-bold text-[#1F3B17] mb-2">
                      {farm.title}
                    </h2>

                    <p className="text-gray-700 text-sm mb-3 line-clamp-3">
                      {farm.description}
                    </p>

                    <div className="text-sm text-gray-800 flex flex-col gap-1 mb-2">
                      <p className="flex items-center gap-2">
                        <MapPin size={16} className="text-[#5E7E3F]" />
                        {farm.location}
                      </p>
                      <p className="flex items-center gap-2">
                        <Leaf size={16} className="text-[#5E7E3F]" />
                        Crop: <span className="font-medium">{farm.crop}</span>
                      </p>
                      <p className="flex items-center gap-2">
                        <TrendingUp size={16} className="text-[#5E7E3F]" />
                        ROI: {Math.round(farm.expected_roi * 100 - 100)}%
                      </p>
                      <p className="flex items-center gap-2">
                        <PackageCheck size={16} className="text-[#5E7E3F]" />
                        Units Left:{" "}
                        <span className="font-semibold text-[#1F3B17]">
                          {unitsLeft}
                        </span>
                      </p>
                    </div>

                    <div className="w-full bg-gray-300/50 rounded-full h-3 mb-2">
                      <div
                        className="bg-[#5E7E3F] h-3 rounded-full transition-all duration-500"
                        style={{
                          width: `${
                            (farm.units_sold / farm.units_available) * 100 || 0
                          }%`,
                        }}
                      />
                    </div>
                    <p className="text-xs text-[#1F3B17] mb-1">
                      Progress: {farm.units_sold || 0} / {farm.units_available} units sold
                    </p>

                    <p className="text-sm text-[#1F3B17] mb-2">
                      <span className="font-semibold">
                        Expected Harvest Return:
                      </span>{" "}
                      ₦{expectedReturn.toLocaleString()}
                    </p>

                    <p className="text-lg font-bold text-[#5E7E3F]">
                      ₦{farm.price_per_unit.toLocaleString()} / unit
                    </p>
                  </div>

                  <button
                    onClick={() => alert(`Investing in: ${farm.title}`)}
                    className="mt-4 bg-[#1F3B17] text-white px-4 py-2 rounded-xl hover:bg-[#345628] transition w-full md:w-max"
                  >
                    Invest Now
                  </button>
                </div>
              </div>
            );
          })}
        </div>
      ) : (
        <div className="text-center text-[#1F3B17]">
          No farms available at the moment. Please check back later.
        </div>
      )}

      <div className="text-center mt-10">
        <p className="text-[#4A2C1D] text-sm">
          Join us in supporting local farmers and securing your food future!
        </p>
      </div>
    </section>
  );
}

export default Farms;
 