step1: clone this project using git clone or by downloading the zip file.
step2: Open source code in any code editor
step3: Run "go mod download" command to install all the packages
step4: Before running this project we need to create database containing following table with fields-
              **Tables**              **Fields**
              1.employees            employee_id, employee_name, role_id, project_id
              2.role                 role_id, role_name
              3.permission           permission_id, permission_name
              4.role_permission      role_id, permission_id
