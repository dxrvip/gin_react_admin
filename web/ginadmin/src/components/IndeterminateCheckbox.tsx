import React, { useState, useEffect } from 'react';
import Box from '@mui/material/Box';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import { PackageItem, FuncItem } from "../utils/dataProvider"
import { useRecordContext } from 'react-admin';





interface childrenProps {
  data: PackageItem;
  handleChange: (event: React.ChangeEvent<HTMLInputElement>, data: PackageItem, childrenIndex: number) => void
}

const Children = (props: childrenProps) => {
  const { data, handleChange } = props
  // 使用 useState 来管理 data




  return (

    <Box sx={{ display: 'flex', flexDirection: 'row', ml: 3 }}>
      {data.func.map((item: FuncItem, index: number) => <FormControlLabel
        key={item.name}
        label={item.alias}
        control={<Checkbox checked={item.active} onChange={(event) => handleChange(event, data, index)} />}
      />
      )}
    </Box>

  )
};




export default function IndeterminateCheckbox(props: { data: PackageItem[], setMenuDatas: (d: any) => void }) {
  const { data, setMenuDatas } = props
  const record = useRecordContext()
  const [packageItem, setPackageItem] = useState<PackageItem[]>((): PackageItem[] => {
    if (record == null && (record as any).menu == null) return data
    return data
  });

  useEffect(() => {
    
    setMenuDatas(packageItem)
  }, [packageItem])


  const handleChange1 = (event: React.ChangeEvent<HTMLInputElement>, index: number) => {
    const thisBool = event.target.checked
    const newPackageItem = packageItem.map((item: PackageItem, i: number) => {
      if (index == i) {
        return {
          ...item,
          func: item.func.map((node: FuncItem) => ({ ...node, active: thisBool }))
        }
      }
      return item
    })

    setPackageItem(newPackageItem)
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>, data: PackageItem, childrenIndex: number) => {
    const thisIndex = packageItem.indexOf(data)
    const newPackageItem = packageItem.map((item: PackageItem, ix: number) => {

      if (thisIndex == ix) {
        return {
          ...item, func: item.func.map((node: FuncItem, i: number) => {
            if (childrenIndex == i) {
              return { ...node, active: event.target.checked }
            }
            return node
          })
        }
      }
      return item
    })
    setPackageItem(newPackageItem)
  };

  const isAllOk = (index: number): boolean => {
    return packageItem[index].func.every(item => item.active == true)
  }




  return (
    <>

      {
        packageItem.map(
          (item: PackageItem, index: number) => (
            <Box key={item.id}>
              <FormControlLabel
                label={item.package}
                control={
                  <Checkbox
                    checked={isAllOk(index)}
                    indeterminate={isAllOk(index) !== true}
                    onChange={(venvt) => handleChange1(venvt, index)}
                  />
                }
              />
              <Children data={item} handleChange={handleChange} />

            </Box>
          ))
      }
    </>
  )

}
