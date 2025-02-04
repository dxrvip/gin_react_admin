import { useGetList } from 'react-admin';
import { Datagrid, TextField } from 'react-admin';

const SubDepartments = ({ id }: any) => {
    // 查询子部门数据
    const { data, isLoading } = useGetList('department', {
        filter: { parent_id: id }, // 根据 parent_id 查询子部门
        pagination: { page: 1, perPage: 100 },
    });

    if (isLoading) return <div>加载中...</div>;

    return (
        <Datagrid data={data} isLoading={isLoading}>
            <TextField source="id" label="部门ID" />
            <TextField source="name" label="部门名称" />
            <TextField source="parent_id" label="父部门" />
        </Datagrid>
    );
};

export default SubDepartments;