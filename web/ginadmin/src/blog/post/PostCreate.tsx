import {
    Create,
    SimpleForm,
    TextInput,
    required,
    maxLength,
    minLength,
    ImageInput,
    ImageField,
    ReferenceInput,
    SelectInput,
} from 'react-admin';
import { RichTextInput } from 'ra-input-rich-text';


const PostCreate = () => (
    <Create title="文章">
        <SimpleForm>
            <TextInput source="title" label="文章标题" validate={[required(), maxLength(100), minLength(2)]} />
            <TextInput source="desc" multiline={true} label="文章摘要" validate={[maxLength(200)]} />
            
            {/* 文章分类 */}
            <ReferenceInput source="cid" reference="category">
                <SelectInput label="文章分类" optionText="name" validate={[required()]} />
            </ReferenceInput>


            <ImageInput source="picture" accept={{ 'image/*': ['.png', '.jpg'] }}>
                <ImageField source="src" title="title" />
            </ImageInput>
            {/* 文章内容 */}
            <RichTextInput fullWidth source="content" label="文章内容" validate={[required(), minLength(10)]} />
        </SimpleForm>
    </Create>
);


export default PostCreate;