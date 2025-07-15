import React from 'react'
import Navbar from '../components/Navbar'
import { Outlet } from 'react-router-dom'

function MainLayout() {
  return (
    <div>
        <Navbar />
        <div className="container mx-auto px-4 py-6">
          {/* Outlet for nested routes */}
          <Outlet />
        </div>
    </div>
  )
}

export default MainLayout