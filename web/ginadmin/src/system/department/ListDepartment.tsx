import {
    List,
    Datagrid,
    TextField,
    EditButton,
    DeleteButton,
    useRecordContext
} from 'react-admin'
import SubDepartments from './SubDepartments';


const DepartmentRow = () => {
    const record = useRecordContext();
    if (!record) return <div>null</div>
    console.log(record);

    return (
        <>
            <tr>
                <td>
                    <TextField source="id" record={record} />
                </td>
                <td>
                    <TextField source="name" record={record} />
                </td>
                <td>
                    <TextField source="parent_id" record={record} />
                </td>
            </tr>

            {/* 展开时显示子部门 */}
            {record && (
                <tr>
                    <td colSpan={3}>
                        <SubDepartments id={record.id} />
                    </td>
                </tr>
            )}
        </>
    );
};


const DepartmentDatagrid = (props: any) => {
    return (
        <Datagrid
            {...props}
            expand={<DepartmentRow />} // 使用自定义展开内容
        >
            <TextField source="id" label="部门ID" />
            <TextField source="name" label="部门名称" />
            <TextField source="parent_id" label="父部门" />

        </Datagrid>
    );
};



function ListDepartment() {
    return (
        <List>
            <DepartmentDatagrid />
        </List>
    );
}

export default ListDepartment;