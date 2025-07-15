import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { Menu, X } from 'lucide-react';
import logo from '../assets/nnefarm.png';

const Navbar = () => {
  const [open, setOpen] = useState(false);

  return (
    <nav className="bg-[#F3E7D5] text-[#4A2C1D] shadow-md fixed w-full z-50">
      <div className="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        {/* Logo and Brand Name */}
        <Link to="/" className="flex items-center gap-2">
          <img src={logo} alt="MyNneFarm Logo" className="h-10 w-10 object-contain" />
          <span className="text-xl font-bold tracking-tight text-[#1F3B17]">MyNneFarm</span>
        </Link>

        {/* Desktop Menu */}
        <div className="hidden md:flex gap-6 items-center text-sm font-medium">
          <Link to="/farms" className="hover:text-[#2F5024]">Farms</Link>
          <Link to="/wallet" className="hover:text-[#2F5024]">Wallet</Link>
          <Link to="/marketplace" className="hover:text-[#2F5024]">Marketplace</Link>
          <Link to="/about" className="hover:text-[#2F5024]">About</Link>
        </div>

        {/* Auth Button */}
        <div className="hidden md:block">
          <button className="bg-[#5E7E3F] text-white px-4 py-2 rounded-lg hover:bg-[#2F5024] transition">
            Sign In
          </button>
        </div>

        {/* Mobile Toggle */}
        <div className="md:hidden">
          <button onClick={() => setOpen(!open)}>
            {open ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
          </button>
        </div>
      </div>

      {/* Mobile Menu */}
      {open && (
        <div className="md:hidden bg-[#F3E7D5] text-[#4A2C1D] px-4 py-2 space-y-3">
          <Link to="/farms" onClick={() => setOpen(false)}>Farms</Link>
          <Link to="/wallet" onClick={() => setOpen(false)}>Wallet</Link>
          <Link to="/marketplace" onClick={() => setOpen(false)}>Marketplace</Link>
          <Link to="/about" onClick={() => setOpen(false)}>About</Link>
          <button className="block w-full text-left mt-2 bg-[#5E7E3F] text-white px-4 py-2 rounded-lg hover:bg-[#2F5024]" onClick={() => setOpen(false)}>
            Sign In
          </button>
        </div>
      )}
    </nav>
  );
};

export default Navbar;
