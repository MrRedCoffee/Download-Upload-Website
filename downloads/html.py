import os
file = open("index.html","w")

text = '''
<style>
a
{
	
}
h1
{
color:white;
text-align: center;
}
.container
{
width: 100vw;
height: 100vh;
background-image: url('lofi.gif');
background-position: center;
background-size:100%;
background-repeat: repeat-y;
background-attachment: fixed;
display: flex;
flex-direction: row;
justify-content: center;
align-items: center;
}
.box
{
width: 300px;
height: 300px;
background: #363636;
border: 3px solid #73AD21;
padding: 10px;
}
.center {
flex-direction: row;
justify-content: center;
align-items: center;
padding: 10px;
}
body
{
margin: 0px;
}

</style>
<html>
'''

file.write(text)
for filename in os.listdir('temp'):
	print(filename)
	text = '''
		<div class="container">
			<div class="box">
				<h1>''' + filename + ''' </h1>
				<div class="center">
					<a href="/download/temp/''' + filename + '''" download>
					<img src="download.png" alt="download" width="50%">
					</a>
				</div>
			</div>
		</div>
	'''
	file.write(text)

text = '''
	</body>
</html>
'''

file.write(text)

file.close()