import React from 'react';
import FormInput from './FormInput';
import { ImageIcon } from 'lucide-react';

function StepThree({ register, errors }) {
  return (
    <div className="grid gap-4">
      <FormInput label="Image URL" name="image_url" icon={<ImageIcon />} register={register} errors={errors} />
      <div>
        <label htmlFor="description" className="block text-sm font-medium text-[#1F3B17] mb-1">
          Farm Description
        </label>
        <textarea
          id="description"
          {...register('description', { required: true })}
          rows="4"
          placeholder="Write details about the farm..."
          className={`w-full border rounded-md p-3 focus:outline-none focus:ring-2 backdrop-blur-md bg-white/30 ${
            errors.description ? 'border-red-500 ring-red-300' : 'border-gray-300 focus:ring-[#5E7E3F]'
          }`}
        />
        {errors.description && <p className="text-red-600 text-xs mt-1">This field is required</p>}
      </div>
    </div>
  );
}

export default StepThree;
