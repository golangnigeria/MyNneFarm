import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { fetchFarms } from '../api/farms';

function Farms() {
  const { data: farms, isLoading, isError } = useQuery({
    queryKey: ['farms'],
    queryFn: fetchFarms,
  });

  const defaultImage = 'https://images.unsplash.com/photo-1580327331400-dc0c0be4a79c?auto=format&fit=crop&w=800&q=60';

  if (isLoading) {
    return <div className="text-center py-10 text-[#1F3B17]">Loading farms...</div>;
  }

  if (isError) {
    return <div className="text-center py-10 text-red-700">Failed to load farms. Please try again later.</div>;
  }

  return (
    <section className="bg-[#F3E7D5] min-h-screen px-4 py-16">
      <h1 className="text-3xl font-bold text-center text-[#1F3B17] mb-10">Available Farms</h1>

      <div className="grid md:grid-cols-3 sm:grid-cols-2 gap-8 max-w-7xl mx-auto">
        {farms.map((farm) => {
          const unitsLeft = farm.units_available - farm.units_sold;
          const roiPercent = Math.round((farm.expected_roi - 1) * 100);

          return (
            <div
              key={farm.id}
              className="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-shadow"
            >
              <img
                src={farm.image_url || defaultImage}
                alt={farm.title}
                className="w-full h-48 object-cover"
                onError={(e) => {
                  e.target.src = defaultImage;
                }}
              />

              <div className="p-5 space-y-2">
                <div className="flex justify-between items-center mb-1">
                  <h2 className="text-lg font-bold text-[#1F3B17]">{farm.title}</h2>
                  {farm.is_active ? (
                    <span className="text-xs px-2 py-1 bg-green-200 text-green-800 rounded-full">Active</span>
                  ) : (
                    <span className="text-xs px-2 py-1 bg-gray-300 text-gray-700 rounded-full">Closed</span>
                  )}
                </div>

                <p className="text-sm text-[#4A2C1D]">
                  {farm.location} | {farm.crop}
                </p>

                <div className="text-sm text-[#4A2C1D] leading-relaxed">
                  <p><strong>₦{farm.price_per_unit.toLocaleString()}</strong> per unit</p>
                  <p>ROI: <strong>{roiPercent}%</strong></p>
                  <p>Units Left: {unitsLeft > 0 ? unitsLeft : 'Sold Out'}</p>
                </div>

                <button
                  disabled={!farm.is_active || unitsLeft <= 0}
                  className={`mt-4 w-full py-2 rounded-lg font-semibold text-white transition-all ${
                    !farm.is_active || unitsLeft <= 0
                      ? 'bg-gray-400 cursor-not-allowed'
                      : 'bg-[#5E7E3F] hover:bg-[#2F5024]'
                  }`}
                >
                  {unitsLeft > 0 ? 'Invest Now' : 'Sold Out'}
                </button>
              </div>
            </div>
          );
        })}
      </div>
    </section>
  );
}

export default Farms;
