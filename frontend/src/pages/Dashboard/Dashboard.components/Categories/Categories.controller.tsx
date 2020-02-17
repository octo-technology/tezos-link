import * as React from 'react'
import { useState } from 'react'

import { CategoriesView } from './Categories.view'
import { CATEGORIES } from './Categories.constants'

export const Categories = () => {
  const [selectedCategory, setSelectedCategory] = useState(0)

  const categoryCallback = (e: any) => {
    console.log(e)
    setSelectedCategory(1)
  }

  return (
    <CategoriesView categories={CATEGORIES} selectedCategory={selectedCategory} categoryCallback={categoryCallback} />
  )
}
