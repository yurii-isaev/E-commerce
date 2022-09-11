import { Button, Form } from 'react-bootstrap'

const RegistrationPage = () => {
  return (
     <Form>
       <Form.Group className="mb-3" controlId="firstname">
         <Form.Label>First name</Form.Label>
         <Form.Control type="firstname" placeholder="Enter first name"/>
       </Form.Group>

       <Form.Group className="mb-3" controlId="lastname">
         <Form.Label>Last name</Form.Label>
         <Form.Control type="lastname" placeholder="Enter last name"/>
       </Form.Group>

       <Form.Group className="mb-3" controlId="email">
         <Form.Label>Email address</Form.Label>
         <Form.Control type="email" placeholder="Enter email address"/>
       </Form.Group>

       <Form.Group className="mb-3" controlId="password">
         <Form.Label>Password</Form.Label>
         <Form.Control type="password" placeholder="Enter password"/>
       </Form.Group>

       <Button variant="primary" type="submit">
         Submit
       </Button>
     </Form>
  )
}

export default RegistrationPage