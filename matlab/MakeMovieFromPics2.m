% fid = fopen('f:\Educ\¿ ”—“» ¿\.œ–Œ√¿\Java\Pressure.txt','r');
% A=fscanf(fid,'%f %f', [2 inf]);
% fclose(fid);

clear all;

writerObj = VideoWriter('d:\MatlabMovie3\YZmovie.avi');
writerObj.FrameRate = 5;
open(writerObj);
%aviobj = avifile('d:\MatlabMovie2\YZmovie.avi','fps',25,'quality',100);

for t=1:111
frame = im2frame(imread(['d:\MatlabMovie3\pics Z\',num2str(t),'.tif']));
writeVideo(writerObj,frame);
end;

for t=1:51
frame = im2frame(imread(['d:\MatlabMovie3\pics Y1\',num2str(t),'.tif']));
writeVideo(writerObj,frame);
end;

for t=1:51
frame = im2frame(imread(['d:\MatlabMovie3\pics Y2\',num2str(t),'.tif']));
writeVideo(writerObj,frame);
end;

for t=1:51
frame = im2frame(imread(['d:\MatlabMovie3\pics Y3\',num2str(t),'.tif']));
writeVideo(writerObj,frame);
end;

close(writerObj);
%aviobj = close(aviobj);